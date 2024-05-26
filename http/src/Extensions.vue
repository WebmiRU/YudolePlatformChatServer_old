<script lang="ts">
import Button from "primevue/button"
import APIService from "./services/APIService"

export default {
  components: {Button},
  data() {
    return {}
  },
  async mounted() {
    console.log('LOG1', await APIService.getExtensions())
  },
  methods: {

  }
}
</script>

<template>
  <h1>Extensions list</h1>
</template>

