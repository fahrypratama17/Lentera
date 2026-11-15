export function OwlArt() {
  return (
    <div
      aria-hidden
      className="relative grid min-h-[280px] place-items-center overflow-hidden rounded-[28px] border border-secondary-light bg-linear-to-br from-secondary-light to-[#cfe6f6] p-6 sm:min-h-[340px] md:p-8"
    >
      <svg viewBox="0 0 320 340" className="h-auto w-full max-w-[330px]" fill="none">
        {/* bintang */}
        <path d="M160 34l7.6 16.4 18 2.4-13.2 12.5 3.3 17.8L160 74.7l-15.7 8.4 3.3-17.8-13.2-12.5 18-2.4z" fill="#0d6684" />
        {/* tumpukan buku kiri */}
        <rect x="30" y="246" width="80" height="16" rx="6" fill="#1f506d" opacity=".85" />
        <rect x="24" y="266" width="92" height="16" rx="6" fill="#0d6684" opacity=".7" />
        <rect x="36" y="286" width="74" height="16" rx="6" fill="#1f506d" opacity=".55" />
        {/* tumpukan buku kanan */}
        <rect x="212" y="246" width="80" height="16" rx="6" fill="#0d6684" opacity=".8" />
        <rect x="204" y="266" width="92" height="16" rx="6" fill="#1f506d" opacity=".75" />
        <rect x="218" y="286" width="74" height="16" rx="6" fill="#0d6684" opacity=".5" />
        {/* badan */}
        <ellipse cx="160" cy="212" rx="66" ry="74" fill="#1f506d" />
        <ellipse cx="160" cy="228" rx="42" ry="50" fill="#e6f3fb" opacity=".92" />
        {/* sayap */}
        <path d="M96 186c-10 26-8 54 6 74 6-26 6-52-6-74z" fill="#0a172b" />
        <path d="M224 186c10 26 8 54-6 74-6-26-6-52 6-74z" fill="#0a172b" />
        {/* kepala */}
        <circle cx="160" cy="126" r="62" fill="#1f506d" />
        {/* mata */}
        <circle cx="136" cy="124" r="21" fill="#ffffff" />
        <circle cx="184" cy="124" r="21" fill="#ffffff" />
        <circle cx="136" cy="124" r="9" fill="#0a172b" />
        <circle cx="184" cy="124" r="9" fill="#0a172b" />
        <circle cx="139" cy="120" r="3" fill="#ffffff" />
        <circle cx="187" cy="120" r="3" fill="#ffffff" />
        {/* paruh */}
        <path d="M160 140l9 12-9 9-9-9z" fill="#0d6684" />
        {/* toga */}
        <path d="M160 62l58 22-58 22-58-22z" fill="#0a172b" />
        <path d="M104 90v18c0 10 25 18 56 18s56-8 56-18V90l-56 21z" fill="#1f506d" />
        <path d="M218 84v30" stroke="#0d6684" strokeWidth="4" strokeLinecap="round" />
        <circle cx="218" cy="118" r="6" fill="#0d6684" />
        {/* kaki */}
        <path d="M138 284v10M182 284v10" stroke="#0d6684" strokeWidth="6" strokeLinecap="round" />
      </svg>
    </div>
  )
}
