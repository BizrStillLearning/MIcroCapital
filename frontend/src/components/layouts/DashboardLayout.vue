<script setup>
import { ref } from 'vue'
import Sidebar from './Sidebar.vue'
import Header from './Header.vue'
import Footer from './Footer.vue'

const isSidebarCollapsed = ref(false)
const isMobileSidebarOpen = ref(false)

const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value
}

const toggleMobileMenu = () => {
  isMobileSidebarOpen.value = true
}

const closeMobileMenu = () => {
  isMobileSidebarOpen.value = false
}
</script>

<template>
  <div class="flex h-screen bg-muted/20" style="font-family: 'DM Sans', sans-serif;">

    <Sidebar
        :isCollapsed="isSidebarCollapsed"
        :isMobileOpen="isMobileSidebarOpen"
        @closeMobile="closeMobileMenu"
    />

    <div class="flex-1 flex flex-col min-w-0 overflow-hidden relative">

      <Header
          :isCollapsed="isSidebarCollapsed"
          @toggleCollapse="toggleSidebar"
          @toggleMobile="toggleMobileMenu"
      />

      <div class="flex-1 overflow-y-auto flex flex-col">

        <main class="flex-1 p-4 sm:p-6 lg:p-8">
          <router-view></router-view>
        </main>

        <Footer />

      </div>

    </div>
  </div>
</template>

