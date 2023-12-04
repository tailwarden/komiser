import React from 'react';
import Image from 'next/image';
import platform from '@utils/providerHelper';
import { ChevronRightIcon } from '@components/icons';

export type CardPrimaryProps = {
  title: string;
  description: string;
  showButton?: boolean;
  showAvatar?: boolean;
  type: 'shadow' | 'stroke';
};

function CardPrimary({
  title,
  description,
  showButton = true,
  showAvatar = true,
  type = 'stroke'
}: CardPrimaryProps) {
  const base = `flex items-center justify-between gap-2 rounded-lg bg-white p-6 border-gray-200 ${
    type === 'shadow' ? 'border-b ' : 'border`'
  }`;

  return (
    <div className={base}>
      <div className="flex gap-4 items-center">
        {showAvatar && (
          <Image
            src={platform.getImgSrc('azure')}
            width={42}
            height={42}
            alt="Purplin thinking"
          />
        )}

        <div>
          <p className="text-lg font-medium text-gray-950">{title}</p>
          <p className="text-gray-700 text-xs">{description}</p>
        </div>
      </div>
      {showButton && <ChevronRightIcon width={16} height={16} />}
    </div>
  );
}

export default CardPrimary;
