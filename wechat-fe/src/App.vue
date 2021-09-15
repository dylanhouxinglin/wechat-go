<template>
    <div id="app">
		<login></login>
		<!-- <welcome></welcome> -->
		<!-- <index></index> -->
    </div>
</template>

<script>
    // import welcome from './components/common/welcome.vue'
	// import index from './index.vue'
	import login from './components/common/login.vue'
    export default {
        name: 'app',
        components: {
            // index,
			login
        },
    }
</script>