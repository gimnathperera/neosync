// export default withPageAuthRequired(async function Home() {
//   return <h1>Hello World!</h1>;
'use client';

import { Button } from '@/components/ui/button';
import NextLink from 'next/link';

// });
export default function Home() {
  return (
    <div>
      <h1>Neosync Overview</h1>
      <NextLink href={'/new/connection'}>
        <Button>New Connection</Button>
      </NextLink>
    </div>
  );
}