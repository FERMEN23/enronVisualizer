<template>
    <div class="h-16 flex items-center justify-between">

      <!-- load all data button -->
      <div class="flex items-center ml-3">
      
        <button @click="getDataFunction()"  title="Reload" class="text-gray-700 px-2 py-1 border border-gray-300 rounded-lg shadow hover:bg-gray-200 transition duration-100">
              
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>

        </button>
      
      </div>

      <!-- search box -->
      <div class="flex items-left bg-white rounded-lg " >

        <div class="w-full py-2 px-4 border border-black-300 rounded-tl-lg rounded-bl-lg" x-data="{ search: '' }">
          
          <input v-model="search" type="search" class="w-full mt-0.5 text-gray-800 focus:outline-none "
                placeholder="search" x-model="search">
        
        </div>

        <div>

          <button @click="modificarValor" type="submit" class="flex  items-center bg-blue-500 justify-center w-12 h-12 text-white rounded-r-lg">
            <svg class="w-5 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </button>

        </div>

      </div>

      <!-- Pagination controller -->
      <div class="px-2 flex items-center space-x-4">

        <span class="text-sm text-gray-500">
          {{ (currentPage - 1) * pageSize +1 }}-{{ (currentPage ) * pageSize }} de {{ length }}
        </span>

        <div class="flex items-center space-x-2">

          <button @click="previousPageFunction" class="bg-gray-200 hover:bg-gray-400 text-gray-700 p-1.5 rounded-lg transition duration-150" title="Newer">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd"></path>
              </svg>
          </button>

          <button @click="nextPageFunction" class="bg-gray-200 hover:bg-gray-400 text-gray-700 p-1.5 rounded-lg transition duration-150" title="Older">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path>
              </svg>
          </button>

        </div>

      </div>
      
    </div>
    
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'EmailListHeader',
  data() {
    return {
      search: '',
    };
  },
  props: {
    currentPage: {
      type: Number,
      required: true,
    },
    pageSize: {
      type: Number,
      required: true,
    },
    length: {
      type: Number,
      required: true,
    },
    getData: {
      type: Function,
      required: true,
    },
    previousPage: {
      type: Function,
      required: true,
    },
    nextPage: {
      type: Function,
      required: true,
    },
    searchTerm: {
      type: String,
      required: true,
    },
  },
  methods: {
    getDataFunction(): void {
      this.getData(1);
    },
    previousPageFunction(): void {
      this.previousPage();
    },
    nextPageFunction(): void {
      this.nextPage();
    },
    modificarValor(): void {
      const nuevoValor: string = this.search;
      this.$emit('valorActualizado', nuevoValor);
    },
  },
  mounted() {
    this.search = this.searchTerm;
  },
});
</script>