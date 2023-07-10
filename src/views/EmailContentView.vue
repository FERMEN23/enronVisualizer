<template>
    <div>
      <div class="flex flex-col md:flex-row items-center m-3">
    <div class="flex-1 px-2">
        <emailContent
            :subject="subject"
            :Initial="initial"
            :Address="from? from.Address : ''"
            :toAsString="toAsString"
            :message="message"
            :dateAndTime="dateAndTime(date)"
            :Cc="Cc"
        ></emailContent>
    </div>
    </div>
    </div>
  </template>
  
<script lang="ts">
import { defineComponent } from 'vue';
import emailContent from '@/components/emailContent.vue';
// import emailContentHeader from '@/components/emailContentHeader.vue';
import axios from 'axios';

interface emailAddress{
  Name: string,
  Address : string
}

interface emailType {
  timestamp: string,
  Body:string,
  Cc: string,
  Content_Transfer_Encoding: string,
  Content_Type: string,
  Date:string,
  ID:number,
  Message_ID: string,
  Mime_Version: string,
  X_FileName: string,
  X_Folder: string,
  X_From: string,
  X_Origin: string,
  X_To: string,
  X_cc: string,
  X_bcc: string,
  from: emailAddress,
  subject: string,
  to: Array<emailAddress>,
}

interface resultType {
  _index: string, 
  _type: string, 
  _id: string, 
  _score: number, 
  _timestamp: string,
  _source : emailType
}

export default defineComponent({
  name: 'EmailcontentView',
  components: {
    emailContent,
  },
  props: ['id'],
  data() {
    return {
      tableData: [],
      pageSize: 20,
      currentPage: 1,
      message: '',
      from: {} as emailAddress,
      subject: '',
      to: [] as emailAddress[],
      toAsString: '',
      Cc: '',
      date: '',
      expanded: false,
      searchTerm: '',
      initial: '',
    };
  },

  methods: {
    getInitials(): string {
      // Obtener las iniciales del nombre
      let initials = '?';
      if (this.from) {
        const names = this.from.Address.split(' ');
        initials = names.map((name: any) => name.charAt(0).toUpperCase()).join('');
      }

      return initials;
    },
    // dateAndTime(fecha: string): string {
    //   const date = new Date(fecha);
    //   const opciones = {
    //     year: 'numeric',
    //     month: 'long',
    //     day: 'numeric',
    //     hour: 'numeric',
    //     minute: 'numeric',
    //     hour12: false,
    //   };
    //   return date.toLocaleDateString('en-US', opciones);
    // },
    dateAndTime(fecha: string): string {
      const date = new Date(fecha);
      const opciones: Intl.DateTimeFormatOptions = { 
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
        hour12: false,
       };
      return date.toLocaleDateString(undefined, opciones);
    },
    async getElement(): Promise<any> {
      try {
        const response = await axios.get('/elementsById/' + this.id);
        this.subject = response.data.hits.hits[0]._source.subject;
        this.from = response.data.hits.hits[0]._source.from;
        this.to = response.data.hits.hits[0]._source.to;
        this.message = response.data.hits.hits[0]._source.Body;
        this.date = response.data.hits.hits[0]._source.Date;
        if (this.to) {
          for (let i = 0; i < this.to.length; i++) {
            this.toAsString += this.to[i].Address;
            if (i < this.to.length - 1) {
              this.toAsString += ', ';
            }
          }
        }

        this.initial = this.getInitials();
      } catch (error) {
        console.error(error);
        return [];
      }
    },
  },
  mounted() {
    this.getElement();
  },
});
</script>
  
  <style>
  /* Component styles... */
  </style>