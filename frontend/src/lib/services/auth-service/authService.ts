import { request } from '$lib/services/api-service/apiService';

/**
 * Logs in a user and stores the token in local storage
 * @param username 
 * @param password 
 * @returns 
 */
export async function login(username: string, password: string): Promise<any> {
    const response = await request('/api/login', 'POST', { username, password });
    if (response.token) {
        localStorage.setItem('token', response.token);
    }
    return response;
}

/**
 * Logs out a user and removes the token from local storage
 * @returns 
 */
export async function logout(): Promise<any> {
    localStorage.removeItem('token');
    return { message: 'Logged out' };
}