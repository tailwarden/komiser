import Image from 'next/image';
import { useContext } from 'react';
import classNames from 'classnames';
import formatNumber from '../../utils/formatNumber';
import GlobalAppContext from '../layout/context/GlobalAppContext';
import StarIcon from '../icons/StarIcon';

type BannerProps = {
  githubStars: number | undefined;
};

function Banner({ githubStars }: BannerProps) {
  const { displayBanner, dismissBanner } = useContext(GlobalAppContext);

  return (
    <div
      className={classNames(
        'top-0 z-10 flex w-full animate-fade-in-down-short items-center justify-center gap-6 bg-gradient-to-br from-primary to-secondary py-3 opacity-0',
        {
          fixed: displayBanner,
          hidden: !displayBanner
        }
      )}
    >
      <span className="text-sm font-medium text-white">
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
          <div className="flex h-full items-center justify-center gap-2 border-l border-white/10 bg-white/10 py-2.5 px-3">
            <StarIcon
              width={16}
              height={16}
              className="group-hover:fill-warning-600 group-hover:text-warning-600"
            />
            {formatNumber(githubStars)}
          </div>
        </a>
      )}

      <button
        onClick={dismissBanner}
        className="absolute right-8 cursor-pointer rounded-lg p-3 text-white transition-colors hover:bg-white/10 active:bg-black-900/10"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M8 16l8-8M16 16L8 8"
          ></path>
        </svg>
      </button>
    </div>
  );
}

export default Banner;
