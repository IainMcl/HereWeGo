<script>
	import "../app.pcss";
    import Header from "./Header.svelte";
    import Footer from "./Footer.svelte";
    import "./styles.css";
</script>

<div class="app">
	<Header></Header>

	<main>
		<slot></slot>
	</main>

	<Footer></Footer>
</div>
