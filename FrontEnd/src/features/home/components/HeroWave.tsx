export function HeroWave() {
  return (
    <div
      aria-hidden="true"
      className="pointer-events-none absolute inset-x-0 bottom-0 z-0 h-[38%] overflow-hidden"
    >
      <svg
        viewBox="0 0 1440 320"
        preserveAspectRatio="none"
        className="absolute inset-x-0 bottom-0 h-full w-full animate-[wave-drift_14s_ease-in-out_infinite_alternate]"
      >
        <path
          fill="#D4EAF5"
          fillOpacity="0.65"
          d="M0,160 C180,40 360,240 560,140 C760,40 980,80 1160,160 C1300,220 1380,120 1440,80 L1440,320 L0,320 Z"
        />
      </svg>
      <svg
        viewBox="0 0 1440 320"
        preserveAspectRatio="none"
        className="absolute inset-x-0 bottom-0 h-[85%] w-full animate-[wave-drift-reverse_18s_ease-in-out_infinite_alternate]"
      >
        <path
          fill="#B5DDEC"
          fillOpacity="0.6"
          d="M0,200 C240,100 420,260 680,170 C900,100 1110,200 1440,130 L1440,320 L0,320 Z"
        />
      </svg>
      <svg
        viewBox="0 0 1440 320"
        preserveAspectRatio="none"
        className="absolute inset-x-0 bottom-0 h-[65%] w-full"
      >
        <path
          fill="#8EC9DF"
          fillOpacity="0.35"
          d="M0,230 C180,170 340,280 560,210 C800,130 1040,280 1240,220 C1340,190 1400,210 1440,180 L1440,320 L0,320 Z"
        />
      </svg>
      <div className="absolute right-[15%] -bottom-25 size-75 rounded-full bg-white/30 blur-[100px]" />
    </div>
  );
}
