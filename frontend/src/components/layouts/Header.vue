<script setup>
import { Menu, PanelLeftClose, PanelLeftOpen, Bell } from '@lucide/vue'

const props = defineProps({
  isCollapsed: Boolean
})

const emit = defineEmits(['toggleCollapse', 'toggleMobile'])
</script>

<template>
  <header class="h-16 bg-card/80 backdrop-blur-md border-b border-border flex items-center justify-between px-4 sm:px-6 sticky top-0 z-10">

    <div class="flex items-center gap-3">
      <button
          @click="emit('toggleMobile')"
          class="md:hidden p-2 -ml-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-xl transition-colors"
      >
        <Menu :size="22" />
      </button>

      <button
          @click="emit('toggleCollapse')"
          class="hidden md:flex p-2 -ml-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-xl transition-colors"
      >
        <PanelLeftOpen v-if="isCollapsed" :size="20" />
        <PanelLeftClose v-else :size="20" />
      </button>

    </div>

    <div class="flex items-center gap-4">
      <button class="relative p-2 text-muted-foreground hover:text-foreground rounded-full hover:bg-muted transition-colors">
        <Bell :size="18" />
        <span class="absolute top-1.5 right-1.5 w-2 h-2 bg-accent rounded-full border border-card"></span>
      </button>

      <div class="flex items-center gap-3 cursor-pointer p-1 pr-2 rounded-full hover:bg-muted transition-colors border border-transparent hover:border-border">
        <img src="https://i.pravatar.cc/150?img=32" alt="User avatar" class="w-8 h-8 rounded-full border border-border" />
        <div class="hidden sm:block text-left">
          <p class="text-sm font-semibold text-foreground leading-none">Grace O.</p>
        </div>
      </div>
    </div>

  </header>
</template>