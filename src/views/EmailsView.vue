<template>
    <HeaderApp></HeaderApp>
      <div v-if="tableData.length==0" class="flex justify-center items-center h-screen">
        <div class="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-gray-900"></div>
      </div>
      <div v-else>
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
      <!-- ToDo: Pasar esto a un componente, que permita modificar el texto search Term -->
      
        <EmailList
        :displayedItems="tableData" 
        >
        </EmailList>
    </div>
    </div> 
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

import HeaderApp from '../components/HeaderApp.vue';
import EmailListHeader from '../components/emailListHeader.vue';
import EmailList from '../components/emailList.vue';
import axios from 'axios';

export default defineComponent({
  name: 'EmailsView',
  components: {
    HeaderApp,
    EmailListHeader,
    EmailList,
  },
  data() {
    return {
      tableData: [] as any[], // Cambia el tipo de tableData a 'any[]'
      pageSize: 20,
      currentPage: 1,
      propValue: 'Valor inicial',
      to: [] as string[], // Cambia el tipo de to a 'string[]'
      totalPagesNum: 0,
      totalEmails: 0,
      searchTerm: '',
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
    // fechaCorta(fecha: string) {
    //   const date = new Date(fecha);
    //   const options = { year: 'numeric', month: 'short', day: 'numeric' };
    //   return date.toLocaleDateString(undefined, options);
    // },
    async getData(pageNumber: number) {
      this.tableData = [];
      const startIndex = (pageNumber - 1) * this.pageSize;
      const endIndex = startIndex + this.pageSize;
      try {
        const response = await axios.get(`/elementsId/${endIndex}`);
        console.log(response)
        this.totalEmails = response.data.hits.total.value;
        this.tableData = response.data.hits.hits;
      } catch (error) {
        console.error(error);
      }
    },
    async getDataByFilter(pageNumber: number) {
      if (this.searchTerm !== '') {
        this.tableData = [];
        const startIndex = (pageNumber - 1) * this.pageSize;
        const endIndex = startIndex + this.pageSize;
        try {
          const response = await axios.get(`/elementsFilter/${this.searchTerm}/${endIndex}`);
          this.totalEmails = response.data.hits.total.value;
          this.tableData = response.data.hits.hits;
        } catch (error) {
          console.error(error);
        }
      } else {
        this.getData(0);
      }
    },
  },
  mounted() {
    this.getData(0);
  },
  // setup() {
  //   const fecha = ref(new Date());
  //   const fechaCorta = ref('');

  //   const cambiarFormatoFecha = () => {
  //     fechaCorta.value = format(fecha.value, 'dd/MM/yyyy');
  //     return fechaCorta.value ;
  //   };

  //   return {
  //     fecha,
  //     fechaCorta,
  //     cambiarFormatoFecha,
  //   };
  // },
});
</script>