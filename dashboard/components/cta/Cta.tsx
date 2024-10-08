import React from 'react';
import Image from 'next/image';

export type CtaProps = {
  title: string;
  description: string;
  action: React.ReactNode;
};

function Cta({ title, description, action }: CtaProps) {
  return (
    <div className="relative flex items-center justify-between gap-2 rounded-lg bg-white px-8 py-6 gradient-border border-[1px]  border-gray-200">
      <div className="flex items-start flex-col gap-4">
        <div className="flex flex-col gap-2">
          <p className="text-gray-950 font-medium text-lg">{title}</p>
          <p className="text-gray-700">{description}</p>
        </div>
        {action}
      </div>
      <Image
        src="/assets/img/purplin/rocket.svg"
        width={123}
        height={128}
        alt="Purplin thinking"
        className="h-[128px] w-[123px]"
      />
    </div>
  );
}

export default Cta;
