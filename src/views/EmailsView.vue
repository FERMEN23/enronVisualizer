<template>
  <div>
      <HeaderApp></HeaderApp>

      <!-- if any email has been load, show a loading spin -->
      <div v-if="tableData.length==0 && !loadedData" class="flex justify-center items-center h-screen">
        <div class="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-gray-900"></div>
      </div>

      <div v-if="tableData.length!=0 && loadedData">
          <div class="flex-1 px-2" x-data="{ checkAll: false, filterMessages: false }">
            <EmailListHeader
                :currentPage= "currentPage"
                :pageSize="pageSize"
                :length="totalEmails"
                :getData = "getData"
                :previousPage = "previousPage"
                :nextPage = "nextPage"
                :searchTerm = "searchTerm"
                @valorActualizado="actualizarValor"
            ></EmailListHeader>

              <EmailList
              :displayedItems="tableData"
              >
              </EmailList>

          </div>

      </div>

      <div v-if="tableData.length==0 && loadedData"
        class="h-screen w-screen flex justify-center">

        <div class="container mt-20 items-center justify-between px-5 text-gray-700">

              <!-- <div class="text-4xl text-black-500 font-dark font-extrabold mb-4">
                Not found
              </div> -->

              <p class="text-xl md:text-xl font-light leading-normal mb-8">
                There is no sender matching with: "{{ searchTerm }}"
              </p>

              <a href="/" class="px-5 inline py-3 text-sm font-medium leading-5 shadow-2xl text-white transition-all duration-400 border border-transparent rounded-lg focus:outline-none bg-gray-600 active:bg-gray-600 hover:bg-gray-800">
                Reload
              </a>

        </div>

      </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import HeaderApp from '../components/HeaderApp.vue';
import EmailListHeader from '../components/emailListHeader.vue';
import EmailList from '../components/emailList.vue';
import axios from 'axios';

import { resultType } from '../types';


export default defineComponent({
  name: 'EmailsView',
  components: {
    HeaderApp,
    EmailListHeader,
    EmailList,
  },
  data() {
    return {
      tableData: [] as resultType[], //list of emails to show
      pageSize: 0,
      currentPage: 1,
      totalEmails: 0,
      searchTerm: '',
      loadedData: false,
    };
  },
  computed: {
    totalPages(): number {
      return Math.ceil(this.totalEmails / this.pageSize);
    },
  },
  methods: {
    actualizarValor(nuevoValor: string) {
      this.searchTerm = nuevoValor;
      if (this.searchTerm === '') {
        this.getData(this.currentPage);
      } else {
        this.getDataByFilter(this.currentPage);
      }
    },
    previousPage() {

      if (this.currentPage > 1) {
        this.currentPage--;
      }

      if (this.searchTerm === '') {
        this.getData(this.currentPage);
      } else {
        this.getDataByFilter(this.currentPage);
      }
    },

    nextPage() {

      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }

      if (this.searchTerm === '') {
        this.getData(this.currentPage);
      } else {
        this.getDataByFilter(this.currentPage);
      }

    },

    async getData(pageNumber: number) {
      
      this.tableData = [];
      this.loadedData = false;

      var startIndex = (pageNumber - 1) * this.pageSize;

      try {
        const response = await axios.get(`/elementsId/${startIndex}`);
        this.totalEmails = response.data.hits.total.value;
        this.tableData = response.data.hits.hits;
        this.loadedData = true;

      } catch (error) {

        console.error(error);

      }
    },
    async getDataByFilter(pageNumber: number) {

      if (this.searchTerm !== '') {

        this.tableData = [];
        this.loadedData = false;
        const startIndex = (pageNumber - 1) * this.pageSize;
        const endIndex = startIndex + this.pageSize;

        try {

          const response = await axios.get(`/elementsFilter/${this.searchTerm}/${endIndex}`);

          this.totalEmails = response.data.hits.total.value;
          this.tableData = response.data.hits.hits;
          this.loadedData = true;

        } catch (error) {

          console.error(error);

        }
      } else {

        this.getData(1);

      }
    },

    async getMaxSize() {
      try {

          const response = await axios.get(`/getMaxSize`);
          this.pageSize= response.data

      } catch (error) {

        console.error(error);

      }
    },

  },
  mounted() {

    this.getData(1);
    this.getMaxSize();

  },
});
</script>