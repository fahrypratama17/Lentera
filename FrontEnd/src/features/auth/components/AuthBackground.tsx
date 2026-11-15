const AuthBackground = () => {
  return (
    <svg
      viewBox="0 0 500 650"
      preserveAspectRatio="none"
      aria-hidden="true"
      className="pointer-events-none absolute inset-0 h-full w-full"
    >
      <defs>
        <linearGradient id="wave-one" x1="0" y1="0" x2="1" y2="1">
          <stop offset="0%" stopColor="#0D6684" />
          <stop offset="100%" stopColor="#70BEDD" />
        </linearGradient>

        <linearGradient id="wave-two" x1="0" y1="0" x2="1" y2="0">
          <stop offset="0%" stopColor="#70BEDD" />
          <stop offset="100%" stopColor="#E6F3FB" />
        </linearGradient>
      </defs>

      <circle cx="440" cy="75" r="180" fill="#E6F3FB" opacity="0.07" />

      <circle cx="30" cy="280" r="120" fill="#70BEDD" opacity="0.06" />

      <path
        d="M0 410 C120 330 250 500 380 420 C430 390 470 370 500 390 L500 650 L0 650 Z"
        fill="url(#wave-one)"
        opacity="0.36"
      />

      <path
        d="M0 490 C140 420 230 570 370 490 C420 460 465 470 500 450 L500 650 L0 650 Z"
        fill="url(#wave-two)"
        opacity="0.22"
      />

      <path
        d="M0 575 C130 515 260 630 390 560 C440 535 470 550 500 525 L500 650 L0 650 Z"
        fill="#E6F3FB"
        opacity="0.12"
      />
    </svg>
  );
};

export default AuthBackground;
