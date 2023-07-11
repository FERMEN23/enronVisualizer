<template>
    <div class=" items-center">

      <HeaderApp></HeaderApp>

      <emailContent
            :subject="subject!=''? subject : '------------ No subject ------------'"
            :Initial="initial"
            :Address="from? from.Address : ''"
            :toAsString="toAsString"
            :message="message"
            :dateAndTime="dateAndTime(date)"
            :Cc="Cc"
        ></emailContent>

  </div>
</template>
  
<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

import emailContent from '@/components/emailContent.vue';
import HeaderApp from '@/components/HeaderApp.vue';

import { formatDate } from '@/utils'
import { emailAddress } from '../types';


export default defineComponent({

  name: 'EmailcontentView',

  components: {
    HeaderApp,
    emailContent,
  },

  props: {
    id: {
      type: String,
      required: true
    },
    
  },
  data() {

    return {

      message: '',
      from: {} as emailAddress,
      subject: '',
      to: [] as emailAddress[],
      toAsString: '',
      Cc: '',
      date: '',
      initial: '',
      
    };
  },

  methods: {

    getInitials(): string {

      let initials = '?';

      if (this.from) {
        const names = this.from.Address.split(' ');
        initials = names.map((name: any) => name.charAt(0).toUpperCase()).join('');
      }

      return initials;
    },

    dateAndTime(fecha: string): string {
      return formatDate(fecha);

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