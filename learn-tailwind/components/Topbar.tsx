import YouTubeIcon from '@mui/icons-material/YouTube';
import InstagramIcon from '@mui/icons-material/Instagram';
import TwitterIcon from '@mui/icons-material/Twitter';

import Link from 'next/link';

export default function Topbar(): JSX.Element {
    return(
        <>
            <div className="p-6 max-w-full mx-auto bg-gray-200 shadow-lg flex items-center font-bebasneue">
                <div className="mr-auto">
                    <ul className="flex items-center space-x-4">
                        <li>
                            <Link className="mr-3" href="/">Home</Link>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <Link className="mr-3" href="#About">About</Link>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <Link className="mr-3" href="#Portfolio">Portfolio</Link>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <Link href="Contact">Contact</Link>
                        </li>
                    </ul>
                </div>
                <div className="ml-auto">
                    <ul className="flex items-center space-x-4">
                        <li>
                            <Link target="_blank" href="https://twitter.com/FrankieLollia"><TwitterIcon className='hover:text-blue-300'/></Link>
                        </li>
                        <li>
                            <Link target="_blank" href="https://www.youtube.com/@frankielollia"><YouTubeIcon className='hover:text-red-500'/></Link>
                        </li>
                        <li>
                            <Link target="_blank" href="https://www.instagram.com/frankielollia/?hl=en"><InstagramIcon className='hover:text-rose-400'/></Link>
                        </li>
                    </ul>
                </div>
            </div>
        </>
    )
}