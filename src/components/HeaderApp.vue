<template>
    <div class="container flex flex-col flex-wrap items-center p-5 mx-auto md:flex-row border-t-2 border-b-2">

      <nav class="flex flex-wrap items-center text-base lg:w-2/5 md:ml-auto">
          <a href="https://www.cs.cmu.edu/~enron/" class="mr-5 font-medium hover:text-gray-900">Enron Home</a>
      </nav>
      
      <div class="flex items-center order-first mb-4 font-medium text-gray-900 lg:order-none lg:w-1/5 title-font lg:items-center lg:justify-center md:mb-0">
        <img src="../assets/logoEnron.png" class="w-[150px] h-[50px]">
      </div>           
      
      <div class="inline-flex items-center h-full ml-5 lg:w-2/5 lg:justify-end lg:ml-0">
          <a href="https://www.cs.cmu.edu/~enron/" class="mr-5 font-medium hover:text-gray-900">Enron Dataset</a>               
      </div>

    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'HeaderApp',
});
</script>