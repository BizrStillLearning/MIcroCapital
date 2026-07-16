<script setup>
import {
  ShieldCheck,
  Lock,
  BellRing,
  Smartphone,
  UserCircle
} from '@lucide/vue'

const userProfile = {
  name: "Grace O.",
  phone: "+62 812 3456 7890",
  joinDate: "Maret 2026",
  isVerified: true,
  verifiedBy: "Agen Budi (Toko Makmur)"
}
</script>

<template>
  <div class="space-y-6 max-w-4xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Pengaturan Akun
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Kelola profil, keamanan PIN, dan preferensi notifikasi Anda.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">

      <div class="md:col-span-1 space-y-6">
        <div class="bg-card border border-border rounded-3xl p-6 text-center shadow-sm">
          <div class="w-20 h-20 bg-primary/10 text-primary rounded-full flex items-center justify-center mx-auto mb-4 border border-primary/20">
            <UserCircle :size="40" />
          </div>
          <h2 class="text-xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">{{ userProfile.name }}</h2>
          <p class="text-sm text-muted-foreground flex items-center justify-center gap-1.5 mt-1" style="font-family: 'DM Mono', monospace;">
            <Smartphone :size="14" /> {{ userProfile.phone }}
          </p>

          <div class="mt-6 pt-6 border-t border-border text-left">
            <p class="text-xs text-muted-foreground mb-2">Status Identitas</p>
            <div
                v-if="userProfile.isVerified"
                class="flex items-start gap-2 text-sm bg-green-50 text-green-700 p-3 rounded-xl border border-green-200"
            >
              <ShieldCheck :size="18" class="flex-shrink-0 mt-0.5" />
              <div>
                <p class="font-bold">Terverifikasi Lokal</p>
                <p class="text-xs opacity-80 mt-0.5">Oleh: {{ userProfile.verifiedBy }}</p>
              </div>
            </div>
            <div v-else class="flex items-center gap-2 text-sm bg-amber-50 text-amber-700 p-3 rounded-xl border border-amber-200">
              <ShieldCheck :size="18" /> Belum Diverifikasi
            </div>
          </div>
        </div>
      </div>

      <div class="md:col-span-2 space-y-6">

        <div class="bg-card border border-border rounded-3xl p-6 shadow-sm">
          <div class="flex items-center gap-3 mb-6">
            <div class="p-2 bg-muted rounded-xl text-foreground"><Lock :size="20" /></div>
            <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Keamanan PIN</h3>
          </div>

          <form @submit.prevent class="space-y-4 max-w-sm">
            <div class="space-y-1.5">
              <label class="text-xs font-bold text-muted-foreground" style="font-family: 'DM Mono', monospace;">PIN SAAT INI</label>
              <input type="password" inputmode="numeric" maxlength="4" placeholder="••••" class="w-full bg-muted/50 border border-border rounded-xl py-2.5 px-4 text-sm tracking-widest focus:outline-none focus:border-primary transition-colors" />
            </div>
            <div class="space-y-1.5">
              <label class="text-xs font-bold text-muted-foreground" style="font-family: 'DM Mono', monospace;">PIN BARU (4-DIGIT)</label>
              <input type="password" inputmode="numeric" maxlength="4" placeholder="••••" class="w-full bg-muted/50 border border-border rounded-xl py-2.5 px-4 text-sm tracking-widest focus:outline-none focus:border-primary transition-colors" />
            </div>
            <button class="bg-primary text-primary-foreground font-bold px-5 py-2.5 rounded-xl hover:opacity-90 transition-opacity text-sm mt-2">
              Perbarui PIN
            </button>
          </form>
        </div>

        <div class="bg-card border border-border rounded-3xl p-6 shadow-sm">
          <div class="flex items-center gap-3 mb-6">
            <div class="p-2 bg-muted rounded-xl text-foreground"><BellRing :size="20" /></div>
            <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Notifikasi SMS / WhatsApp</h3>
          </div>

          <div class="space-y-4">
            <label class="flex items-center justify-between p-4 bg-muted/30 rounded-2xl cursor-pointer hover:bg-muted/50 transition-colors border border-border/50">
              <div>
                <p class="font-bold text-sm text-foreground">Pengingat Tagihan & Iuran</p>
                <p class="text-xs text-muted-foreground mt-0.5">Kirim pesan H-3 sebelum tanggal jatuh tempo cicilan.</p>
              </div>
              <input type="checkbox" class="w-5 h-5 accent-primary rounded cursor-pointer" checked />
            </label>

            <label class="flex items-center justify-between p-4 bg-muted/30 rounded-2xl cursor-pointer hover:bg-muted/50 transition-colors border border-border/50">
              <div>
                <p class="font-bold text-sm text-foreground">Aktivitas Dompet</p>
                <p class="text-xs text-muted-foreground mt-0.5">Kirim pesan setiap ada uang masuk atau keluar dari dompet Anda.</p>
              </div>
              <input type="checkbox" class="w-5 h-5 accent-primary rounded cursor-pointer" checked />
            </label>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>