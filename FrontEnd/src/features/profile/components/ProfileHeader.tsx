import { UserRound } from "lucide-react";

import ProfileWave from "./ProfileWave";

interface ProfileHeaderProps {
  name: string;
  email: string;
  initials: string;
}

const ProfileHeader = ({ name, email, initials }: ProfileHeaderProps) => {
  return (
    <section className="relative overflow-hidden rounded-3xl border border-[#DCE6ED] bg-white p-6 shadow-sm md:p-8">
      <ProfileWave />

      <div className="relative z-10 flex flex-col items-start gap-5 sm:flex-row sm:items-center">
        <div className="bg-primary flex size-22 shrink-0 items-center justify-center rounded-full border-4 border-white text-2xl font-bold text-white shadow-md">
          {initials || <UserRound className="size-10" />}
        </div>

        <div className="min-w-0">
          <h2 className="text-2xl-bold text-primary-text md:text-3xl-bold mt-1 wrap-break-word">
            {name || "Pengguna StudyTrack"}
          </h2>
          <p className="text-sm-medium text-secondary-text mt-2 break-all">
            {email || "Email tidak tersedia"}
          </p>
        </div>
      </div>
    </section>
  );
};

export default ProfileHeader;
