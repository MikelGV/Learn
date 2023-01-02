import Head from "next/head";

const siteTitle: string = 'Tailwind test website'

export default function Seo() {
    return(
        <Head>
            <meta
            name="description"
            content="Frankie Taylor Scott website"
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
        
    )
}