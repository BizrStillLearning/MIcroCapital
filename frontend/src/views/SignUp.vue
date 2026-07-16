<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft, ArrowRight, User, Smartphone, Lock } from '@lucide/vue'
import { useAuthStore } from '../stores/auth.js'

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
  name: '',
  phone: '',
  pin: ''
})

const handleSignUp = () => {
  console.log('Register attempt:', form)
  authStore.login()
  router.push('/dashboard')
}
</script>

<template>
  <div class="min-h-screen bg-background flex" style="font-family: 'DM Sans', sans-serif;">

    <div class="hidden md:flex md:w-1/2 relative bg-primary items-center justify-center overflow-hidden">
      <div class="absolute inset-0">
        <img src="https://images.unsplash.com/photo-1542884748-2b87b36f6b90?w=1200&h=1080&fit=crop&auto=format" alt="Community" class="w-full h-full object-cover opacity-30 mix-blend-multiply" />
        <div class="absolute inset-0 bg-gradient-to-t from-primary via-primary/80 to-transparent"></div>
      </div>
      <div class="relative z-10 p-12 max-w-lg text-primary-foreground">
        <div class="flex items-center gap-2.5 mb-8">
          <div class="w-10 h-10 rounded-lg bg-white flex items-center justify-center">
            <span class="text-primary font-bold text-xl" style="font-family: 'Fraunces', serif;">U</span>
          </div>
          <span class="font-bold text-2xl text-white" style="font-family: 'Fraunces', serif;">Umoja</span>
        </div>
        <h1 class="text-4xl lg:text-5xl font-bold mb-6 leading-tight" style="font-family: 'Fraunces', serif;">Mulai Perjalanan Anda.</h1>
        <p class="text-lg opacity-80 leading-relaxed">Bergabunglah dengan ribuan anggota lainnya membangun masa depan melalui inklusi keuangan tanpa batas.</p>
      </div>
    </div>

    <div class="w-full md:w-1/2 flex items-center justify-center p-5 sm:p-12 relative">
      <div class="w-full max-w-md">
        <button @click="router.push('/')" class="absolute top-8 left-5 sm:left-12 text-sm text-muted-foreground hover:text-foreground inline-flex items-center gap-2 transition-colors">
          <ArrowLeft :size="16" /> Kembali
        </button>

        <div class="mb-10 mt-8 md:mt-0">
          <h2 class="text-3xl font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">Buat Akun Baru</h2>
          <p class="text-muted-foreground text-sm">Hanya butuh nomor telepon untuk memulai.</p>
        </div>

        <form @submit.prevent="handleSignUp" class="space-y-5">
          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Nama Lengkap</label>
            <div class="relative">
              <User class="absolute left-3.5 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
              <input v-model="form.name" type="text" placeholder="Contoh: Budi Santoso" class="w-full bg-card border border-border rounded-xl py-3 pl-10 pr-4 text-sm focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" required />
            </div>
          </div>

          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Nomor Telepon</label>
            <div class="relative">
              <Smartphone class="absolute left-3.5 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
              <input v-model="form.phone" type="tel" placeholder="+62 812 3456 7890" class="w-full bg-card border border-border rounded-xl py-3 pl-10 pr-4 text-sm focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" required />
            </div>
          </div>

          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Buat PIN 4-Digit</label>
            <div class="relative">
              <Lock class="absolute left-3.5 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
              <input v-model="form.pin" type="password" inputmode="numeric" maxlength="4" placeholder="••••" class="w-full bg-card border border-border rounded-xl py-3 pl-10 pr-4 text-sm tracking-widest focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" required />
            </div>
          </div>

          <button type="submit" class="w-full bg-primary text-primary-foreground font-semibold py-3.5 rounded-xl hover:opacity-90 hover:shadow-lg transition-all flex items-center justify-center gap-2 mt-2">
            Buat Akun <ArrowRight :size="18" />
          </button>
        </form>

        <div class="mt-8 text-center text-sm text-muted-foreground">
          Sudah punya akun?
          <button @click="router.push('/signin')" class="text-primary font-bold hover:underline ml-1" style="font-family: 'DM Mono', monospace;">Masuk di sini</button>
        </div>
      </div>
    </div>
  </div>
</template>