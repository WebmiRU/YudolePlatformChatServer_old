<script setup lang="ts">
import store from "../store.ts";
</script>

<script lang="ts">
import APIService from "../services/APIService"

export default {
  components: {},
  data() {
    return {
      model: {payload: {params: {}}},
    }
  },
  async mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'modules.index'}},
      {label: 'Modules', route: {name: 'modules.index'}},
      {label: this.$route.params.id, route: {name: 'modules.id', params: {id: this.$route.params.id}}},
    ]

    this.model = await APIService.getModulesId(this.$route.params.id)
  },
  methods: {

  }
}
</script>

<template>
  {{model.payload}}
  <h1>Module {{$route.params.id}} params</h1>

  <TabView>
<!--    <TabPanel v-for="tab in Object.keys(model.payload.params)" :key="tab" :header="tab">-->
<!--      <p class="m-0">{{ tab }}</p>-->
<!--    </TabPanel>-->
  </TabView>
</template>

