<script setup>
import { useRouter, useRoute } from 'vue-router'
import { LogOut, X } from '@lucide/vue'
import { getAsideMenu } from '../../config/menuSidebar.js'

const sidebarLinks = getAsideMenu('member')
const router = useRouter()
const route = useRoute()

const props = defineProps({
  isCollapsed: Boolean,
  isMobileOpen: Boolean
})

const emit = defineEmits(['closeMobile'])

const navigateTo = (path) => {
  router.push(path)
  emit('closeMobile')
}
</script>

<template>
  <aside
      class="hidden md:flex flex-col bg-card border-r border-border transition-all duration-300 ease-in-out z-20"
      :class="isCollapsed ? 'w-20' : 'w-64'"
  >
    <div class="h-16 flex items-center justify-center border-b border-border transition-all">
      <div class="w-8 h-8 rounded-lg bg-primary flex items-center justify-center flex-shrink-0">
        <span class="text-primary-foreground font-bold text-sm" style="font-family: 'Fraunces', serif;">U</span>
      </div>
      <span
          v-if="!isCollapsed"
          class="font-bold text-lg text-foreground tracking-tight ml-2 truncate"
          style="font-family: 'Fraunces', serif;"
      >
        Umoja
      </span>
    </div>

    <nav class="flex-1 px-3 py-6 space-y-2 overflow-y-auto overflow-x-hidden">
      <button
          v-for="link in sidebarLinks"
          :key="link.name"
          @click="navigateTo(link.path)"
          class="w-full flex items-center px-3 py-3 rounded-xl cursor-pointer transition-colors focus:outline-none"
          :class="[
          route.path === link.path ? 'bg-primary/10 text-primary font-semibold' : 'text-muted-foreground hover:bg-muted hover:text-foreground font-medium',
          isCollapsed ? 'justify-center' : 'gap-3'
        ]"
          :title="isCollapsed ? link.name : ''"
      >
        <component :is="link.icon" :size="20" class="flex-shrink-0" />
        <span v-if="!isCollapsed" class="text-sm truncate text-left">{{ link.name }}</span>
      </button>
    </nav>

    <div class="p-4 border-t border-border">
      <button
          class="flex items-center px-3 py-3 w-full rounded-xl text-red-500 hover:bg-red-50 transition-colors font-medium"
          :class="isCollapsed ? 'justify-center' : 'gap-3'"
          :title="isCollapsed ? 'Keluar' : ''"
      >
        <LogOut :size="20" class="flex-shrink-0" />
        <span v-if="!isCollapsed" class="text-sm truncate">Keluar</span>
      </button>
    </div>
  </aside>

  <Transition
      enter-active-class="transition-opacity duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-300"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
  >
    <div v-if="isMobileOpen" class="fixed inset-0 z-40 md:hidden flex">
      <div class="fixed inset-0 bg-black/60 backdrop-blur-sm" @click="emit('closeMobile')"></div>

      <aside class="relative w-[80%] max-w-sm bg-card h-full flex flex-col shadow-2xl transition-transform">
        <div class="h-16 flex items-center justify-between px-5 border-b border-border">
          <span class="font-bold text-lg text-foreground" style="font-family: 'Fraunces', serif;">Menu Dasbor</span>
          <button @click="emit('closeMobile')" class="p-2 -mr-2 text-muted-foreground hover:text-foreground rounded-xl hover:bg-muted transition-colors">
            <X :size="20" />
          </button>
        </div>

        <nav class="flex-1 px-4 py-6 space-y-2 overflow-y-auto">
          <button
              v-for="link in sidebarLinks"
              :key="link.name"
              @click="navigateTo(link.path)"
              class="w-full flex items-center gap-3 px-3 py-3 rounded-xl cursor-pointer transition-colors focus:outline-none"
              :class="route.path === link.path ? 'bg-primary/10 text-primary font-semibold' : 'text-muted-foreground hover:bg-muted font-medium'"
          >
            <component :is="link.icon" :size="20" />
            <span class="text-left">{{ link.name }}</span>
          </button>
        </nav>
      </aside>
    </div>
  </Transition>
</template>

