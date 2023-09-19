import React from 'react';
import Image from 'next/image';

import ProviderCls, { Provider } from '../../utils/providerHelper';

function PurplinCloud({ provider }: { provider: Provider }) {
  return (
    <div className="relative inline-block">
      <Image
        src="/assets/img/others/onboarding-cloud.svg"
        alt="Onboarding cloud"
        width={500}
        height={120}
      />
      <div className="absolute top-[53%] left-[48%] -translate-x-1/2 -translate-y-1/2 transform rounded-full">
        <Image
          src={ProviderCls.providerImg(provider) as string}
          alt={`${provider} Logo`}
          className="rounded-full shadow-md"
          width={95}
          height={95}
        />
      </div>
    </div>
  );
}

export default PurplinCloud;
