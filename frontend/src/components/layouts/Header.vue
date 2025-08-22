<template>
  <header class="mx-auto py-3 flex items-center justify-between">
    <a href="/" class="flex items-center gap-3">
      <div class="w-10 h-10 bg-gradient-to-br from-primary to-primary-light rounded-lg flex items-center justify-center text-white font-extrabold">SI</div>
      <div>
        <p class="text-xl font-semibold">StocksInfo</p>
      </div>
    </a>

    <nav class="hidden md:flex items-center gap-6 text-sm text-slate-600">
      <a href="/" class="hover:text-slate-900">Features</a>
      <a href="/companies" class="hover:text-slate-900">Companies</a>
      <a href="/recommendations" class="">
        <button class="bg-primary text-white hover:bg-primary-light">Recommendations</button>
      </a>
    </nav>


    <button class="md:hidden p-2 rounded-md bg-white shadow"><SearchIcon/></button>
  </header>
</template>
<script setup lang="ts">

import { SearchIcon } from 'lucide-vue-next'

</script>