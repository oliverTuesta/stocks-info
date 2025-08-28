<template>
  <main>
    <section class="relative pt-8">
      <div class="bg-primary text-white p-6 rounded-md">
        <h1>Welcome to stocks-info</h1>
        <p>The best place to see the <b>real worth</b> of your stocks</p>
        <img class="hero-image" src="/spigi.png" alt="Oliver Tuesta cartoon drawing"/>
        <button class="bg-white hover:bg-gray-300 text-black mt-4">Let's go</button>
      </div>
    </section>

    <section id="features" class="container mt-16 mx-auto lg:px-6">
      <h2>Features</h2>
      <p>Packed with experts recommendations powered with AI tools for new investors</p>

      <div class="mt-6 grid grid-cols-1 sm:grid-cols-2  gap-6">
        <div v-for="f in features" :key="f.title" class="p-5 rounded-lg bg-gray-100 flex flex-row gap-5">
          <div class="w-20 h-20 flex-none flex items-center justify-center my-auto rounded bg-white">
            <component :is="f.icon" class="w-8 h-8 text-black" />
          </div>
          <div>
            <h4 class="font-bold">{{ f.title }}</h4>
            <p class="text-sm mt-2">{{ f.desc }}</p>
          </div>
        </div>
      </div>
    </section>
    <section class="container mt-16 mx-auto lg:px-6">
      <HotCompanies/>
    </section>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { UserPenIcon, PanelBottomDashedIcon, SquirrelIcon } from 'lucide-vue-next'
import HotCompanies from '@/components/HotCompanies.vue'

const features = ref([{
  title: 'Investor fit',
  desc: 'Matches a stock to passive or aggressive strategies based on volatility and fundamentals.',
  icon: UserPenIcon
}, {
  title: 'Horizon guidance',
  desc: 'Short vs long term indicators and suggested holding periods.',
  icon: PanelBottomDashedIcon
}, {
  title: 'Explainable rationale',
  desc: 'Short, human friendly reasons and risk points for each suggestion.',
  icon: SquirrelIcon
}, {
  title: 'Relevant information',
  desc: 'AI powered analysis and recommendations for each investor profile.',
  icon: UserPenIcon
}])
</script>

<style scoped>
.hero-image {
  height: 220px;
  position: absolute;
  top: -4px;
  right: 30px;
}

@media (max-width: 665px) {
  .hero-image {
    display: none;
  }
}
</style>