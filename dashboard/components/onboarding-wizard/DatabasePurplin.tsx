import Image from 'next/image';

interface DatabasePurplinProps {
  database: string;
}

function DatabasePurplin({ database }: DatabasePurplinProps) {
  return (
    <div className="relative">
      <div className="-mb-14 -ml-12 flex w-24 items-center justify-center rounded-3xl bg-cyan-200 p-4">
        <Image
          src={`/assets/img/database/${
            database === 'postgres' ? 'postgresql' : 'sqlite'
          }.svg`}
          alt={`${database} Logo`}
          className="h-16 w-16"
          width={0}
          height={0}
        />
      </div>
      <Image
        src="/assets/img/purplin/dashboard.svg"
        alt={`Purplin Dashboard Logo`}
        className="h-60 w-60"
        width={0}
        height={0}
      />
    </div>
  );
}

export default DatabasePurplin;
