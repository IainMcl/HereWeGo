<script>
	import "../app.pcss";
	import Header from "./Header.svelte";
	import Footer from "./Footer.svelte";
	import "./styles.css";
</script>

<div class="app flex flex-col h-screen justify-between bg-background">
	<Header></Header>

	<main class="mb-auto">
		<slot />
	</main>

	<Footer></Footer>
</div>
