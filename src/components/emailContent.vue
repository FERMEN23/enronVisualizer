<template>
    <div>

        <div class="flex items-center justify-between flex-1 mb-3 py-2 px-2">

            <a @click="navigateToHome" href="#" class="flex  text-gray-700 px-2 py-1 space-x-0.5 border border-gray-300 rounded-lg shadow hover:bg-gray-200 transition duration-100" title="Back">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M9.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L7.414 9H15a1 1 0 110 2H7.414l2.293 2.293a1 1 0 010 1.414z" clip-rule="evenodd"></path>
                </svg>

                <span class="text-sm font-bold">Back</span>                                   
            </a>

            <div class="flex-1">

                <h4 class="text-lg text-gray-800 font-bold ">
                    {{ subject }}
                </h4>

            </div>

        </div>
    
        <div class="flex p-2  justify-center">

            <div class=" w-1/4 flex items-center justify-center w-10 h-10 rounded-full bg-gray-500 text-white">
                <span class="text-xl font-bold">{{ Initial }}</span>
            </div>

            <div class="w-1/2">

                <div class="flex">

                    <div class="flex flex-col items-start ml-2">

                        <div class="text-xs text-gray-400 text-left">
                            From: {{ Address }}
                        </div>  

                        <div class="text-xs text-gray-400 max-w-[550px] text-left" >
                            
                            <div :class="{'truncate': !expanded, 'line-clamp-3': !expanded}">
                                To: {{ toAsString }}
                            </div>
                            
                            <button @click="expanded = !expanded" class="text-gray-500 hover:underline mt-1">
                                {{ expanded ? 'Show less' : 'Show more...' }}
                            </button>

                        </div>

                        <div>
                        
                            <div v-if="Cc!=''" class="text-sm text-gray-400 max-w-[600px]" >
                                
                                <div  :class="{'truncate': !expanded, 'line-clamp-3': !expanded}">
                                    Cc: {{ Cc }}
                                </div>

                                <button @click="expandedCc = !expandedCc" class="text-gray-500 hover:underline mt-1">
                                    {{ expandedCc ? 'Show more' : 'Show less...' }}
                                </button>
                                
                            </div>

                        </div>

                    </div>   
                            
                </div>

                <div class="py-6 pl-2 text-gray-700 whitespace-pre-wrap text-justify" :style="{ width: '750px' }">
                    {{ message }}
                </div>  

            </div>   

            <div class=" w-1/4 flex  w-30 h-30 rounded-full ">
                <span class="text-sm text-gray-500">{{ dateAndTime }}</span>
            </div>  

        </div>
        
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Router, useRouter } from 'vue-router';

export default defineComponent({

    name: 'EmailContent',

    data() {

        return {
            expanded: false,
            expandedCc: false 
        
        }
    },
    props: {
        subject: {
            type: String,
            required: true
        },
        Initial: {
            type: String,
            required: true
        },
        Address: {
            type: String,
            required: true
        },
        toAsString: {
            type: String,
            required: true
        },
        message: {
            type: String,
            required: true
        },
        dateAndTime: {
            type: String,
            required: true
        },
        Cc: {
            type: String,
            required: true
        },
    },
    setup() {
        const router: Router = useRouter();

        const navigateToHome = (): void => {
            router.push('/');
            
        };

        return {
            navigateToHome,
            
        };
  },
}   
)
</script>