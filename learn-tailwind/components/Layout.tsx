import Head from "next/head";
import styles from '../styles/Home.module.css'

const siteTitle: string = 'Tailwind test website'

export default function Layout({children, home}: {children: any, home: boolean}) {
    return(
        <div>
            <Head>
                <link rel="icon" href="/favicon.ico" />
                <meta
                name="description"
                content="Learn how to use tailwind"
                />
                <meta
                property="og:image"
                content={`https://og-image.vercel.app/${encodeURI(
                    siteTitle,
                )}.png?theme=light&md=0&fontSize=75px&images=https%3A%2F%2Fassets.vercel.com%2Fimage%2Fupload%2Ffront%2Fassets%2Fdesign%2Fnextjs-black-logo.svg`}
                />
                <meta name="og:title" content={siteTitle} />
                <meta name="twitter:card" content="summary_large_image" />
            </Head>



            <header>
                {home ? 
                (
                <>
                    <h1 className='text-3xl font-bold underline'>Hello World!</h1>
                </>
                ): (
                <>
                    <h2 className="text-3xl font-bold">Hi World!</h2>
                </>
                )}
            </header>
            <main>{children}</main>
            <footer>
                <h1>this is footer</h1>
            </footer>
        </div>
    )
}