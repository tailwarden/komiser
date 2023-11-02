import React from 'react';
import Image from 'next/image';

import Avatar from '@components/avatar/Avatar';
import { Provider } from '../../utils/providerHelper';

function PurplinCloud({ provider }: { provider: Provider }) {
  return (
    <div className="relative inline-block">
      <Image
        src="/assets/img/others/onboarding-cloud.svg"
        alt="Onboarding cloud"
        width={500}
        height={120}
      />
      <div className="absolute left-[48%] top-[53%] -translate-x-1/2 -translate-y-1/2 transform rounded-full">
        <Avatar avatarName={provider} size={96} />
      </div>
    </div>
  );
}

export default PurplinCloud;
