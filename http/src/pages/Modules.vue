<script lang="ts" setup>
import store from "../store"
</script>

<script lang="ts">
import APIService from "../services/APIService"

export default {
  computed: {},
  components: {},
  data() {
    return {
      model: {payload: {}},
    }
  },
  async mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'index'}},
      {label: 'Modules', route: {name: 'modules.index'}}
    ]

    this.model = await APIService.getModules()
  },
  methods: {
    f1(v) {
      console.log(v)
    }
  }
}
</script>

<template>
  <h1>Modules list</h1>

  <br/>
  <br/>

  <DataTable :value="model.payload" tableStyle="min-width: 50rem">
    <Column field="name" header="Name"></Column>

    <Column header="On/Off">
      <template #body="row">
        <InputSwitch v-model="checked" />
      </template>
    </Column>

    <Column header="Active">
      <template #body="row">
        <Badge v-if="row.data.is_active" severity="success">Active</Badge>
        <Badge v-else severity="danger">Inactive</Badge>
      </template>
    </Column>

    <Column header="Config">
      <template #body="row">
        <RouterLink :to="{name: 'modules.id', params: {id: row.index}}">
          <Button label="Config" severity="secondary"/>
        </RouterLink>
      </template>
    </Column>
  </DataTable>

</template>

