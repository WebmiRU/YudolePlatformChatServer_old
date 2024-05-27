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
  <h1>Module {{$route.params.id}} params</h1>

  <TabView>
    <TabPanel v-for="key in Object.keys(model.payload.params)" :key="key" :header="key">
      <div v-for="field in model.payload.params[key]" class="field flex flex-column gap-1 mb-5">
          <label v-if="field.label.length">{{ field.label }}</label>
          <InputText v-if="field.type == 'string'" v-model="value" />
          <InputNumber v-else-if="field.type == 'number'" v-model="value" />

          <small v-if="field.description.length">{{field.description}}</small>
      </div>

    </TabPanel>
  </TabView>
</template>

