<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
</script>

<script lang="ts">
import Button from "primevue/button";

export default {
  components: {Button},
  data() {
    return {

    }
  },
  mounted() {

  },
  methods: {
    f1() {
      alert('F1')
    }
  }
  }
</script>

<template>
  <Button label="Primary" raised @click="f1" />
  <Button label="Secondary" severity="secondary" raised />
  <Button label="Success" severity="success" raised />
  <Button label="Info" severity="info" raised />
  <Button label="Warning" severity="warning" raised />
  <Button label="Help" severity="help" raised />
  <Button label="Danger" severity="danger" raised />
  <Button label="Contrast" severity="contrast" raised />

</template>
