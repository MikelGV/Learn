import YouTubeIcon from '@mui/icons-material/YouTube';
import InstagramIcon from '@mui/icons-material/Instagram';
import TwitterIcon from '@mui/icons-material/Twitter';
import Link from 'next/link';


export default function Footer() {
    return(
        <>
            <footer className="text-center bg-gray-200">
                <div className="p-6 max-w-full mx-auto bg-gray-200 shadow-lg font-bebasneue flex items-center text-center">
                    <div className="mr-auto">
                        <ul>
                            <li>
                                <Link href="/contact">Contact</Link>
                            </li>
                        </ul>
                    </div>
                    <div className="ml-auto">
                        <ul className="flex items-center space-x-4">
                            <li>
                                <a target="_blank" href="https://twitter.com/FrankieLollia"><TwitterIcon className='hover:text-blue-300 text-lg'/></a>
                            </li>
                            <li>
                                <a target="_blank" href="https://www.youtube.com/@frankielollia"><YouTubeIcon className='hover:text-red-500 text-lg'/></a>
                            </li>
                            <li>
                                <a target="_blank" href="https://www.instagram.com/frankielollia/?hl=en"><InstagramIcon className='hover:text-rose-400 text-lg'/></a>
                            </li>
                        </ul>
                    </div>
                </div>
                <div className="p-2 max-w-full mx-auto bg-gray-300 shadow-lg font-bebasneue text-center">
                    <div className="flex justify-center md:mr-auto">
                        <h1>Made By <Link href="https://www.mikelgaldos.com" target="_blank">Mikel Galdos</Link></h1>
                    </div>
                </div>
            </footer>
        </>
    )
}