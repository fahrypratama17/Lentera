import {
  UserPlus, Mail, BookOpen, Eye,
  Target, UserCheck, BarChart3,
} from "lucide-react"

export const NAV_LINKS = [
  { label: "Cara menggunakan", href: "#cara-menggunakan" },
  { label: "Manfaat", href: "#manfaat" },
  { label: "Tentang", href: "#tentang" },
]

export const STEPS = [
  { n: 1, icon: UserPlus, title: "Buat cohort",
    text: "Isi nama, deskripsi, dan periode belajar. Pembuat otomatis menjadi Mentor.",
    variant: "tint" as const },
  { n: 2, icon: Mail, title: "Undang anggota",
    text: "Masukkan email atau tautan undangan. Tentukan role Mentor atau Member.",
    variant: "plain" as const },
  { n: 3, icon: BookOpen, title: "Susun materi",
    text: "Mentor menambah materi berurutan dengan panduan, referensi, dan deadline.",
    variant: "tint" as const },
  { n: 4, icon: Eye, title: "Belajar & review",
    text: "Member mengumpulkan hasil. Mentor menyetujui atau meminta revisi.",
    variant: "plain" as const },
]

export const BENEFITS = [
  { icon: Target, title: "Pemantauan tanpa bertanya",
    text: "Metrik anggota × materi menunjukkan siapa tertinggal pada materi mana, dalam satu pandangan." },
  { icon: UserCheck, title: "Review yang tidak menumpuk",
    text: "Setiap pengumpulan masuk antrean review. Mentor memutuskan dengan konteks, bukan dari ingatan." },
  { icon: BarChart3, title: "Keterlambatan terdeteksi dini",
    text: "Deadline jadi acuan, bukan pengingat. Progres yang lewat deadline ditandai terpisah." },
]

export const STATS = [
  { label: "Role", num: "2", desc: "Mentor & Member, per cohort." },
  { label: "Status", num: "4", desc: "Belum mulai sampai Selesai." },
  { label: "Matriks", num: "1", desc: "Seluruh progres dalam satu tampilan." },
]
