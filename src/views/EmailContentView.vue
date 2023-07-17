<template>
    <div class=" items-center">
      <HeaderApp></HeaderApp>
      <div v-if="errorOcurred">
          <ErrorComponent
          :codeError="errorObject.code"
          :nameError="errorObject.name"
          :messageError="errorObject.message"
          ></ErrorComponent>
      </div>

      <div v-else>
        <emailContent
            :emailSubject="emailSubject!=''? emailSubject : '------------ No subject ------------'"
            :fromInitial="fromInitial"
            :fromAddress="emailFrom? emailFrom.Address : ''"
            :emailToAsString="emailToAsString"
            :emailMessage="emailMessage"
            :dateAndTime="dateAndTime(emailDate)"
            :Cc="emailCc"
        ></emailContent>
      </div>
  </div>
</template>
  
<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

import emailContent from '@/components/emailContent.vue';
import HeaderApp from '@/components/HeaderApp.vue';
import ErrorComponent from '@/components/errorComponent.vue';

import { formatDate } from '@/utils'
import { emailAddress } from '../types';


export default defineComponent({

  name: 'EmailcontentView',

  components: {
    HeaderApp,
    emailContent,
    ErrorComponent,
  },

  props: {
    id: {
      type: String,
      required: true
    },
    
  },
  data() {
    return {
      emailMessage: '',
      emailFrom: {} as emailAddress,
      emailSubject: '',
      emailTo: [] as emailAddress[],
      emailToAsString: '',
      emailCc: '',
      emailDate: '',
      fromInitial: '',
      errorOcurred: false,
      errorObject: {name: '', message: '', code: ''}
    };
  },

  methods: {
    getInitials(): string {
      let initials = '?';
      if (this.emailFrom) {
        const names = this.emailFrom.Address.split(' ');
        initials = names.map((name: any) => name.charAt(0).toUpperCase()).join('');
      }

      return initials;
    },

    dateAndTime(fecha: string): string {
      return formatDate(fecha);
    },

    async getEmailById(): Promise<any> {
      this.errorOcurred = false;
      try {
        const response = await axios.get('/v1/emailById/' + this.id);
        if (response.data.hits.total.value >0) {
          this.emailSubject = response.data.hits.hits[0]._source.subject;
          this.emailFrom = response.data.hits.hits[0]._source.from;
          this.emailTo = response.data.hits.hits[0]._source.to;
          this.emailMessage = response.data.hits.hits[0]._source.Body;
          this.emailDate = response.data.hits.hits[0]._source.Date;

          if (this.emailTo) {
            for (let i = 0; i < this.emailTo.length; i++) {
              this.emailToAsString += this.emailTo[i].Address;
              if (i < this.emailTo.length - 1) {
                this.emailToAsString += ', ';
              }
            }
          }
          this.fromInitial = this.getInitials();
        } else {
          this.errorOcurred = true;
          this.errorObject = {
            name: 'Email no found',
            message: 'Email with _id: ' + this.id + ' not founded',
            code: '400'
          }
        }
        
      } catch (error: any) {
        console.error("Error", error);
        this.errorOcurred = true;
        this.errorObject.name = error.name
        this.errorObject.message = "Sorry " + error.message
        this.errorObject.code = error   
      }
    },
  },

  mounted() {
    this.getEmailById();

  },
});
</script>