<template>
  <div>
    <HeaderApp></HeaderApp>
    <!-- if any email has been load, show a loading spin -->
    <div v-if="loadSpinner" class="flex justify-center items-center h-screen">
      <div class="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-gray-900"></div>
    </div>

    <div v-if="displayItems">
      <div class="flex-1 px-2" x-data="{ checkAll: false, filterMessages: false }">
        <EmailListHeader 
            :currentPage="currentPage" 
            :paginationSize="paginationSize" 
            :totalEmails="totalEmails" 
            :getEmailsPagination="getEmailsPagination"
            :previousPage="previousPage" 
            :nextPage="nextPage" 
            :searchTerm="searchTerm" 
            @valueUpdated="updateValue"></EmailListHeader>
        <EmailList :emailstoDisplay="emailsList"></EmailList>
      </div>
    </div>

    <div v-if="itemNotFounded" class="h-screen w-screen flex justify-center">
      <div class="container mt-20 items-center justify-between px-5 text-gray-700">
        <p class="text-xl md:text-xl font-light leading-normal mb-8">
          There is no sender matching with: "{{ searchTerm }}"
        </p>
        <a href="/"
          class="px-5 inline py-3 text-sm font-medium leading-5 shadow-2xl text-white transition-all duration-400 border border-transparent rounded-lg focus:outline-none bg-gray-600 active:bg-gray-600 hover:bg-gray-800">
          Reload
        </a>
      </div>
    </div>

    <div v-if="errorOcurred">
      <ErrorComponent
      :codeError="errorObject.code"
      :nameError="errorObject.name"
      :messageError="errorObject.message"
      ></ErrorComponent>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import HeaderApp from '../components/HeaderApp.vue';
import EmailListHeader from '../components/emailListHeader.vue';
import EmailList from '../components/emailList.vue';
import ErrorComponent from '@/components/errorComponent.vue';
import axios from 'axios';

import { resultType } from '../types';


export default defineComponent({
  name: 'EmailsListView',
  components: {
    HeaderApp,
    EmailListHeader,
    EmailList,
    ErrorComponent,
  },
  data() {
    return {
      emailsList: [] as resultType[], //list of emails to show
      paginationSize: 0,
      currentPage: 1,
      totalEmails: 0,
      searchTerm: '',
      loadedData: false,
      errorOcurred: false,
      errorObject: {name: '', message: '', code: ''}
    };
  },
  computed: {
    totalPages(): number {
      return Math.ceil(this.totalEmails / this.paginationSize);
    },
    loadSpinner(): boolean{
      return this.emailsList.length == 0 && !this.loadedData && !this.errorOcurred
    },
    displayItems(): boolean {
      return this.emailsList.length != 0 && this.loadedData && !this.errorOcurred
    },
    itemNotFounded(): boolean {
      return this.emailsList.length == 0 && this.loadedData && !this.errorOcurred
    }
  },
  methods: {
    updateValue(newValue: string) {
      this.searchTerm = newValue;
      if (this.searchTerm === '') {
        this.getEmailsPagination(this.currentPage);
      } else {
        this.getEmailsByFilter(this.currentPage);
      }
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }

      if (this.searchTerm === '') {
        this.getEmailsPagination(this.currentPage);
      } else {
        this.getEmailsByFilter(this.currentPage);
      }

    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }

      if (this.searchTerm === '') {
        this.getEmailsPagination(this.currentPage);
      } else {
        this.getEmailsByFilter(this.currentPage);
      }

    },
    async getEmailsPagination(pageNumber: number) {
      this.emailsList = [];
      this.loadedData = false;
      this.errorOcurred = false;
      this.errorObject= {name: '', message: '', code: ''}

      const startIndex = (pageNumber - 1) * this.paginationSize + 1;

      try {
        const response = await axios.get(`/emailsPagination/${startIndex}`);

        this.emailsList = response.data.hits.hits;
        this.loadedData = true;

        this.getTotalEmails()

      } catch (error: any) {
        this.errorOcurred = true;
        this.errorObject.name = error.name
        this.errorObject.message = "Sorry " + error.message +" ocurred"
        this.errorObject.code = error.code
        console.error("Error", error);
      }
    },
    async getEmailsByFilter(pageNumber: number) {
      if (this.searchTerm !== '') {
        this.emailsList = [];
        this.loadedData = false;
        this.errorObject= {name: '', message: '', code: ''}
        const startIndex = (pageNumber - 1) * this.paginationSize;
     
        try {
          const response = await axios.get(`/emailsFilter/${this.searchTerm}/${startIndex}`);

          this.totalEmails = response.data.hits.total.value;
          this.emailsList = response.data.hits.hits;
          this.loadedData = true;

        } catch (error: any) {
            this.errorOcurred = true;
            this.errorObject.name = error.name
            this.errorObject.message = "Sorry an " + error.message +" ocurred"
            this.errorObject.code = error.code
       // console.error("Error",this.errorObject);
      }
      } else {
        this.getEmailsPagination(1);
      }
    },

    async getMaxResult() {
      this.errorObject= {name: '', message: '', code: ''}
      try {
        const response = await axios.get(`/getMaxResultVariable`);
        this.paginationSize = response.data

      } catch (error: any) {
        this.errorOcurred = true;
        this.errorObject.name = error.name
        this.errorObject.message = "Sorry an " + error.message +" ocurred"
        this.errorObject.code = error.code
       // console.error("Error",this.errorObject);
      }
    },

    async getTotalEmails() {
      this.errorObject= {name: '', message: '', code: ''}
      try {
        const response = await axios.get(`/getallEmails`);
        console.log(response)
        this.totalEmails = response.data.hits.total.value;

      } catch (error: any) {
        this.errorOcurred = true;
        this.errorObject.name = error.name
        this.errorObject.message = "Sorry an " + error.message +" ocurred"
        this.errorObject.code = error.code
       // console.error("Error",this.errorObject);
      }
    },

  },
  mounted() {
    this.getEmailsPagination(1);
    this.getMaxResult();

  },
});
</script>