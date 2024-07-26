import { StarIcon } from '@components/icons';
import formatNumber from '@utils/formatNumber';
import Image from 'next/image';
import { BannerProps } from './Banner';

const base: BannerProps = {
  children: 'Banner Content',
  displayBanner: true,
  dismissBanner: () => {}
};

const primary: BannerProps = {
  children: (
    <>
      <span className="text-sm font-medium">
        Support Komiser by giving us a star on GitHub.
      </span>

      <a
        href="https://github.com/tailwarden/komiser"
        target="_blank"
        rel="noreferrer"
        className="group flex items-center gap-3 rounded border-[1.5px] border-white pl-4 text-sm text-white transition-colors hover:bg-white/10"
      >
        <Image
          src="/assets/img/others/github-white.svg"
          width="18"
          height="16"
          alt="Github logo"
        />
        <span>Star Komiser</span>
        <div className="flex h-full items-center justify-center gap-2 border-l border-white/10 bg-white/10 px-3 py-2.5">
          <StarIcon
            width={16}
            height={16}
            className="group-hover:fill-orange-400 group-hover:text-orange-400"
          />
          {formatNumber(100000)}
        </div>
      </a>
    </>
  ),
  displayBanner: true,
  dismissBanner: () => {}
};

const secondary: BannerProps = {
  children: (
    <>
      <span className="text-sm font-medium">
        Support Komiser by giving us a star on GitHub.
      </span>

      <a
        href="https://github.com/tailwarden/komiser"
        target="_blank"
        rel="noreferrer"
        className="group flex items-center gap-3 rounded border-[1.5px] border-darkcyan-500 text-darkcyan-500 pl-4 text-sm transition-colors hover:bg-black/10"
      >
        <Image
          src="/assets/img/others/github-black.svg"
          width="18"
          height="16"
          alt="Github logo"
        />
        <span>Star Komiser</span>
        <div className="flex h-full items-center justify-center gap-2 border-l border-black/10 bg-black/10 px-3 py-2.5">
          <StarIcon
            width={16}
            height={16}
            className="group-hover:fill-orange-400 group-hover:text-orange-400"
          />
          {formatNumber(100000)}
        </div>
      </a>
    </>
  ),
  displayBanner: true,
  dismissBanner: () => {},
  style: 'secondary'
};

const mockBannerProps = {
  base,
  primary,
  secondary
};

export default mockBannerProps;
