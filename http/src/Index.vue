<script lang="ts" setup>
import store from "./store.ts"
</script>

<script lang="ts">
export default {
  mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'index'}},
    ]
  },
  methods: {}
}
</script>

<template>
  <h1>Index page</h1>
</template>

