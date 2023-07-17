<template>
    <div class="w-full h-screen flex flex-col lg:flex-row items-center justify-center space-y-16 lg:space-y-0 space-x-8 2xl:space-x-0">
    <div class="w-full lg:w-1/2 flex flex-col items-center justify-center lg:px-2 xl:px-0 text-center">
        <p class="text-2xl md:text-4xl lg:text-6xl font-bold tracking-wider text-gray-300">{{ codeError }}</p>
        <p class="text-2xl md:text-4xl lg:text-6xl font-bold tracking-wider text-gray-300 mt-2">{{ nameError }}</p>
        <p class="text-lg md:text-xl lg:text-2xl text-gray-500 my-12">{{ messageError }}</p>
        <a href="/" class="flex items-center space-x-2 bg-blue-600 hover:bg-blue-700 text-gray-100 px-4 py-2 rounded transition duration-150" title="Return Home"> 
            <span>Retry</span>
        </a>
    </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'ErrorComponent',
    props: {
        codeError: {
            type: String,
            required: true   
        },
        nameError: {
            type: String,
            required: true   
        },
        messageError: {
            type: String,
            required: true   
        }
    }
    
});
</script>