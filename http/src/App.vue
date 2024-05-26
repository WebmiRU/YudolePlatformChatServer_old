<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from 'vue-router';

const router = useRouter();

const items = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    command: () => {
      router.push({name: 'home'})
    }
  },
  {
    label: 'Extensions',
    icon: 'pi pi-box',
    command: () => {
      router.push({name: 'extensions'})
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
import Menubar from 'primevue/menubar';
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
  <Menubar :model="items" />

  <RouterLink :to="{name: 'route1'}">Link 1</RouterLink>
  <RouterLink :to="{name: 'route2'}">Link 2</RouterLink>

  <strong>Current route path:</strong> {{ $route.fullPath }}

  <RouterView />


<!--  <Button label="Primary" raised @click="f1" />-->
<!--  <Button label="Secondary" severity="secondary" raised />-->
<!--  <Button label="Success" severity="success" raised />-->
<!--  <Button label="Info" severity="info" raised />-->
<!--  <Button label="Warning" severity="warning" raised />-->
<!--  <Button label="Help" severity="help" raised />-->
<!--  <Button label="Danger" severity="danger" raised />-->
<!--  <Button label="Contrast" severity="contrast" raised />-->

</template>
