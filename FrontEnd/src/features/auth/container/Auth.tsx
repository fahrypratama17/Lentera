import AuthCard from "@/features/auth/components/AuthCard";

const Auth = () => {
  return (
    <main className="bg-surface relative flex min-h-dvh w-full items-center justify-center overflow-hidden px-4 py-8 sm:px-6">
      <div
        aria-hidden="true"
        className="bg-secondary-light pointer-events-none absolute -top-40 -left-40 size-96 rounded-full blur-3xl"
      />

      <div
        aria-hidden="true"
        className="bg-secondary-light pointer-events-none absolute -right-40 -bottom-40 size-96 rounded-full blur-3xl"
      />

      <div className="relative z-10 flex w-full justify-center">
        <AuthCard />
      </div>
    </main>
  );
};

export default Auth;
