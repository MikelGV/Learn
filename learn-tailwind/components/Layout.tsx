import Head from "next/head";
import Topbar from "./Topbar";

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
                {/* <div className="p-6 max-w-full mx-auto bg-gray-200 shadow-lg flex items-center space-x-4">
                    <div className="ml-auto">
                        <ul className="flex items-center space-x-4">
                            <li>
                                <a href="/">Home</a>
                            </li>
                            <li>
                                <a href="#About">About</a>
                            </li>
                            <li>
                                <a href="#Portfolio">Portfolio</a>
                            </li>
                            <li>
                                <a href="Contact">Contact</a>
                            </li>
                        </ul>
                    </div>
                </div> */}
                <Topbar />

            </header>
            <main>{children}</main>
            <footer>
                <h1 className="text-2xl font-light">this is footer</h1>
            </footer>
        </div>
    )
}