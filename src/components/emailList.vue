<template>
  <div class="bg-gray-100 mb-6">
    <ul>
      <li v-for="item in emailstoDisplay" :key="item._id" @click="navigateToView(item._id);"
        class="flex items-center border-y hover:bg-gray-200 px-2">
        <div x-data="{ messageHover: false }" @mouseover="messageHover = true" @mouseleave="messageHover = false"
          class="w-full flex items-center justify-between p-1 my-1 cursor-pointer">
          <div class="flex items-center">
            <div class="flex mr-4 ml-1 space-x-1">
              <button>
                <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M20 12L4 12M4 12L12 20M4 12L12 4M20 12L12 4L12 20"></path>
                </svg>
              </button>
            </div>
            <span class="w-56 pr-2 truncate">
              {{ item._source.from ? item._source.from.Address : " " }}
            </span>
            <span class="w-96 truncate">
              {{ item._source.subject !== '' ? item._source.subject : '------------ No subject ------------' }}
            </span>
            <span class="mx-1">-</span>
            <span class="w-96 text-gray-600 text-sm truncate">
              {{ item._source.Body ? item._source.Body : item._source.Body }}
            </span>
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

import { shortFormatDate } from '@/utils'
import { resultType } from '../types';

export default defineComponent({
  name: 'EmailList',
  data() {
    return {
      messageHover: false,
    };
  },
  props: {
    emailstoDisplay: {
      type: Array as PropType<resultType[]>,
      required: true
    },

  },
  methods: {
    ShortDate(date: string): string {
      return shortFormatDate(date)
    },
  },
  setup() {
    const router: Router = useRouter();
    const navigateToView = (id: string): void => {
      router.push(`/EmailContent/${id}`);
    };
    return {
      navigateToView
    };
  }
});
</script>