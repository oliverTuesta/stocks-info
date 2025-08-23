<template>
  <header class="mx-auto py-3 flex items-center justify-between">
    <router-link to="/" class="flex items-center gap-3">
      <div class="w-10 h-10 bg-gradient-to-br from-primary to-primary-light rounded-lg flex items-center justify-center text-white font-extrabold">SI</div>
      <div>
        <p class="text-xl font-semibold">StocksInfo</p>
      </div>
    </router-link>

    <nav class="hidden md:flex items-center gap-6 text-sm text-slate-600">
      <router-link to="/" class="hover:text-slate-900">Features</router-link>
      <router-link to="/companies" class="hover:text-slate-900">Companies</router-link>
      <router-link to="/recommendations" class="">
        <button class="bg-primary text-white hover:bg-primary-light">Recommendations</button>
      </router-link>
    </nav>


    <button class="md:hidden p-2 rounded-md bg-white shadow"><SearchIcon/></button>
  </header>
</template>
<script setup lang="ts">

import { SearchIcon } from 'lucide-vue-next'

</script>