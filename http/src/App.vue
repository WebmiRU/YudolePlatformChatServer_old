<script setup lang="ts">
import {ref} from "vue";
import {useRouter} from "vue-router";
import store from "./store"

const router = useRouter();

const items = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    command: () => {
      router.push({name: 'index'})
    }
  },
  {
    label: 'Modules',
    icon: 'pi pi-box',
    command: () => {
      router.push({name: 'modules.index'})
    }
  },
  // {
  //   label: 'Features',
  //   icon: 'pi pi-star'
  // },
  // {
  //   label: 'Projects',
  //   icon: 'pi pi-search',
  //   items: [
  //     {
  //       label: 'Components',
  //       icon: 'pi pi-bolt'
  //     },
  //     {
  //       label: 'Blocks',
  //       icon: 'pi pi-server'
  //     },
  //     {
  //       label: 'UI Kit',
  //       icon: 'pi pi-pencil'
  //     },
  //     {
  //       label: 'Templates',
  //       icon: 'pi pi-palette',
  //       items: [
  //         {
  //           label: 'Apollo',
  //           icon: 'pi pi-palette'
  //         },
  //         {
  //           label: 'Ultima',
  //           icon: 'pi pi-palette'
  //         }
  //       ]
  //     }
  //   ]
  // },
  // {
  //   label: 'Contact',
  //   icon: 'pi pi-envelope'
  // }
]);
</script>

<script lang="ts">

export default {
  components: {},
  data() {
    return {}
  },
  mounted() {

  },
  methods: {}
}
</script>

<template>
  <Menubar :model="items"/>

  <Breadcrumb v-if="store.breadcrumbs.length > 1" :model="store.breadcrumbs">
    <template #item="{ item, props }">
      <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
        <a :href="href" v-bind="props.action" @click="navigate">
          <span :class="[item.icon, 'text-color']"/>
          <span class="text-primary font-semibold">{{ item.label }}</span>
        </a>
      </router-link>
      <a v-else :href="item.url" :target="item.target" v-bind="props.action">
        <span class="text-color">{{ item.label }}</span>
      </a>
    </template>
  </Breadcrumb>

  <!--  <RouterLink :to="{name: 'route1'}">Link 1</RouterLink>-->
  <!--  <RouterLink :to="{name: 'route2'}">Link 2</RouterLink>-->

  <!--  <strong>Current route path:</strong> {{ $route.fullPath }}-->

  <RouterView/>


  <!--  <Button label="Primary" raised @click="f1" />-->
  <!--  <Button label="Secondary" severity="secondary" raised />-->
  <!--  <Button label="Success" severity="success" raised />-->
  <!--  <Button label="Info" severity="info" raised />-->
  <!--  <Button label="Warning" severity="warning" raised />-->
  <!--  <Button label="Help" severity="help" raised />-->
  <!--  <Button label="Danger" severity="danger" raised />-->
  <!--  <Button label="Contrast" severity="contrast" raised />-->

</template>
