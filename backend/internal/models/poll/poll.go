package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Poll interface {
	CreatePoll(db *sqlx.DB) error
	UpdatePoll(db *sqlx.DB) error
}

type PollOption interface {
	CreatePollOption(db *sqlx.DB) error
	DeletePollOption(db *sqlx.DB) error
	VoteFor(db *sqlx.DB, userId int64) error
	VoteAgainst(db *sqlx.DB, userId int64) error
}

type NamedPoll struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Options     []PollOption `json:"options"`
	CreatedById int64        `json:"created_by"  db:"created_by"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

type NamedPollOption struct {
	ID        int64     `json:"id" db:"id"`
	PollId    int64     `json:"poll_id" db:"poll_id"`
	Name      string    `json:"name"`
	VoteCount int       `json:"vote_count" db:"vote_count"`
	VotedBy   []int64   `json:"voted_by"` // TODO: Not implemented yet
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type vote struct {
	UserId       int64 `db:"user_id"`
	PollOptionId int64 `db:"poll_option_id"`
}

func (p *NamedPoll) CreatePoll(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		INSERT INTO polls 
			(title, description, created_by) 
			VALUES ($1, $2, $3);
			`, p.Title, p.Description, p.CreatedById)

	tx.Commit()
	return nil
}

func GetPollById(db *sqlx.DB, id int64) (*NamedPoll, error) {
	p := &NamedPoll{}
	err := db.Get(p,
		`SELECT 
		id
		, title
		, description
		, created_by
		, created_at
		, updated_at
	FROM polls WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	po := []NamedPollOption{}
	err = db.Select(&po,
		`SELECT
			id
			, poll_id
			, name
			, vote_count
			, created_at
			, updated_at
		FROM poll_options WHERE poll_id = $1`, id)
	if err != nil {
		return nil, err
	}
	po_ids := []int64{}
	for _, v := range po {
		po_ids = append(po_ids, v.ID)
	}
	var votes []vote
	err = db.Select(&votes,
		`SELECT user_id, poll_option_id FROM votes WHERE poll_option_id = ANY($1)`, po_ids)
	if err != nil {
		return nil, err
	}

	p.Options = make([]PollOption, len(po))
	for i, v := range po {
		for _, vote := range votes {
			if vote.PollOptionId == v.ID {
				v.VoteCount++
				v.VotedBy = append(v.VotedBy, vote.UserId)
			}
		}
		val := v
		p.Options[i] = &val
	}
	return p, nil
}

func (p *NamedPoll) UpdatePoll(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		UPDATE polls
		SET title = $1, description = $2, updated_at = now()
		WHERE id = $3;
		`, p.Title, p.Description, p.ID)
	tx.Commit()
	return nil
}

func (po *NamedPollOption) CreatePollOption(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		INSERT INTO poll_options 
			(poll_id, name) 
			VALUES ($1, $2);
			`, po.PollId, po.Name)

	tx.Commit()
	return nil
}

func (po *NamedPollOption) DeletePollOption(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec(`
		DELETE FROM poll_options
		WHERE id = $1;
		`, po.ID)

	tx.Commit()
	return nil
}

func (po *NamedPollOption) VoteFor(db *sqlx.DB, userId int64) error {
	tx := db.MustBegin()
	tx.MustExec(`
		UPDATE poll_options 
		SET votes = votes + 1, voted_by = array_append(voted_by, $1)
		, updated_at = now()
		WHERE id = $2;
		`, userId, po.ID)

	tx.Commit()
	return nil
}

func VoteFor(db *sqlx.DB, userId, pollOptionId int64) error {
	tx := db.MustBegin()
	tx.MustExec(`
		INSERT INTO votes (user_id, poll_option_id) VALUES ($1, $2);
		`, userId, pollOptionId)

	tx.Commit()
	return nil
}

func RemoveVote(db *sqlx.DB, userId, pollOptionId int64) error {
	tx := db.MustBegin()
	tx.MustExec(`
		DELETE FROM votes WHERE user_id = $1 AND poll_option_id = $2;
		`, userId, pollOptionId)

	tx.Commit()
	return nil
}

func (po *NamedPollOption) VoteAgainst(db *sqlx.DB, userId int64) error {
	tx := db.MustBegin()
	tx.MustExec(`
		UPDATE poll_options 
		SET votes = votes - 1, voted_by = array_remove(voted_by, $1) 
		, updated_at = now()
		WHERE id = $2;
		`, userId, po.ID)

	tx.Commit()
	return nil
}
