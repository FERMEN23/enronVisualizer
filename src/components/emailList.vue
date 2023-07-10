<template>
  <div class="bg-gray-100 mb-6">
    <ul>
      <li v-for="item in displayedItems" :key="item._id" @click="navigateToView(item._id);" class="flex items-center border-y hover:bg-gray-200 px-2">
        <div 
          x-data="{ messageHover: false }"  
          @mouseover="messageHover = true" 
          @mouseleave="messageHover = false" 
          class="w-full flex items-center justify-between p-1 my-1 cursor-pointer">
          <div class="flex items-center">
            <div class="flex items-center mr-4 ml-1 space-x-1">
              <button>
                <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12L4 12M4 12L12 20M4 12L12 4M20 12L12 4L12 20"></path>
                </svg>
              </button>
              <!-- <span class="w-56 pr-2 truncate">{{ item._source.ID ? item._source.ID : " " }}</span> -->
            </div>
            <span class="w-56 pr-2 truncate">{{ item._source.from ? item._source.from.Address : " " }}</span>
            <span v-if="item._source.subject !== ''" class="w-96 truncate">{{ item._source.subject }}</span>
            <span v-else class="w-96 truncate text-gray-600">No subject</span>
            <span class="mx-1">-</span>
            <span class="w-96 text-gray-600 text-sm truncate">{{ item._source.Body ? item._source.Body : item._source.Body }}</span>
          </div>
          <div class="w-32 flex items-center justify-end">
            <span x-show="!messageHover" class="text-sm text-gray-500">
              {{ ShortDate(item._source.Date) }}
            </span>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { Router, useRouter } from 'vue-router';

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
  from: any,
  subject: string,
  to: Array<any>,
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
  name: 'EmailList',
  data() {
    return {
      messageHover: false
    };
  },
  props: {
    displayedItems: {
      type: Array as PropType<resultType[]>,
      required: true
    }
  },
  methods: {
    ShortDate(fecha: string): string {
      const date = new Date(fecha);
      const opciones: Intl.DateTimeFormatOptions =  { year: 'numeric', month: 'short', day: 'numeric' };
      return date.toLocaleDateString(undefined, opciones);
    },
  },
  setup() {
    const router: Router = useRouter();

    const navigateToView = (id: string): void => {
      console.log(router, id);
      router.push(`/EmailContent/${id}`);
      
    };
    return {
      navigateToView
    };
  }
});
</script>