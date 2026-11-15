"use client"

import Link from "next/link"
import { Users } from "lucide-react"
import Image from "next/image"

export function Logo({ onNavigate }: { onNavigate?: () => void }) {
  return (
    <Link
      href="#top"
      onClick={onNavigate}
      aria-label="Lentera - Learning Tracker"
      className="flex shrink-0 items-center gap-2 rounded-full w-1/4
        focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
    >
      <Image src={"/logo.webp"} alt="Logo-Navbar" width={500} height={500} className="w-[10%]"/>
      <span className="flex flex-col leading-tight">
        <span className="text-[20px] font-extrabold tracking-tight text-primary">Lentera</span>
      </span>
    </Link>
  )
}
