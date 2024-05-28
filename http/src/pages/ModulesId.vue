<script setup lang="ts">
// @ts-ignore
import store from "../store.ts";
</script>

<script lang="ts">
import APIService from "../services/APIService"

export default {
  components: {},
  data() {
    return {
      model: null,
    }
  },
  async mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'modules.index'}},
      {label: 'Modules', route: {name: 'modules.index'}},
      {label: this.$route.params.id, route: {name: 'modules.id', params: {id: this.$route.params.id}}},
    ]

    this.model = await APIService.modulesIdGet(this.$route.params.id)
  },
  methods: {
    save(id: string, payload: object) {
      APIService.modulesIdPut(id, payload)
    }
  }
}
</script>

<template>
  <h1>Module {{ $route.params.id }} params</h1>

  <TabView>
    <TabPanel v-if="model?.payload" v-for="tabKey in Object.keys(model.payload.params)" :key="tabKey" :header="tabKey">
      <div v-for="(field, fieldKey) in model.payload.params[tabKey]" class="field flex flex-column gap-1 mb-5">
        <label v-if="field.label.length">{{ field.label }}</label>
        <InputText
            v-if="field.type == 'string'"
            v-model="model.payload.params[tabKey][fieldKey].value"
            :placeholder="model.payload.params[tabKey][fieldKey].placeholder"
        />
        <InputNumber
            v-else-if="field.type == 'number'"
            v-model="model.payload.params[tabKey][fieldKey].value"
            :placeholder="model.payload.params[tabKey][fieldKey].placeholder"
        />

        <small v-if="field.description.length">{{ field.description }}</small>
      </div>

      <div  class="field gap-1 mb-2">
        <Button label="Save" severity="success" @click="save($route.params.id, this.model)"/>
      </div>

    </TabPanel>
  </TabView>
</template>

