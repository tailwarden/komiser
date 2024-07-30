import Banner from '@components/banner/Banner';
import { StarIcon } from '@components/icons';
import GlobalAppContext from '@components/layout/context/GlobalAppContext';
import Image from 'next/image';
import { useContext } from 'react';
import formatNumber from '../../../../utils/formatNumber';

type GithubBannerProps = {
  githubStars: number | undefined;
};

function GithubBanner({ githubStars }: GithubBannerProps) {
  const { displayBanner, dismissBanner } = useContext(GlobalAppContext);

  return (
    <Banner displayBanner={displayBanner} dismissBanner={dismissBanner}>
      <span className="text-sm font-medium">
        Support Komiser by giving us a star on GitHub.
      </span>

      {githubStars && (
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
            {formatNumber(githubStars)}
          </div>
        </a>
      )}
    </Banner>
  );
}

export default GithubBanner;
