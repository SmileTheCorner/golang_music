<template>
    <router-view v-slot="{Component}">
     <transition name="fade">
        <component :is="Component" class="container"></component>
     </transition>
    </router-view>
</template>
<style scoped lang="scss">
.fade-enter-from{
    opacity: 0;
}
.fade-enter-active{
    transition: all 2s;
}
.fade-enter-to{
    opacity: 1;
}

</style>