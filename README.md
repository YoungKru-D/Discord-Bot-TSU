Discord is an instant messaging and VoIP social platform which allows communication through voice calls, video calls, text messaging, and media and files. Communication can be private or take place in virtual communities called "servers". A server is a collection of persistent chat rooms and voice channels which can be accessed via invite links. Discord runs on Windows, macOS, Android, iOS, iPadOS, Linux, and in web browsers. As of 2021, the service has over 350 million registered users and over 150 million monthly active users. It is primarily used by gamers, although the share of users interested in other topics is growing.

In December 2016, the company introduced its GameBridge API, which allows game developers to directly integrate with Discord within games. In December 2017, Discord added a software development kit that allows developers to integrate their games with the service, called "rich presence". This integration is commonly used to allow players to join each other's games through Discord or to display information about a player's game progression in their Discord profile.

Bots are community-made tools to automate tasks. When installed by server owners, they may aid in moderation, host mini games, and perform myriad of other automated tasks. As of 2021, there are around 430,000 total bots active in estimated 30% of all servers. Discord provides official bot APIs which allow custom elements such as dropdowns and buttons. In spring 2022, Discord released an official "app directory" where server owners can add bots to their servers in-Discord. The Verge described bots as an "important part of Discord".

Project Planning

Functional Requirements:

-User Interaction: Allow users to interact with the bot through text commands.
-Moderation: Implement moderation features to manage users, messages, and server settings.
-Information Retrieval: Retrieve data, such as server information, and user profiles.
-Customization: Provide options for server owners or admin to customize bot behavior or features.

Non-Functional Requirements:

-Performance: Ensure the bot operates efficiently and responds promptly to commands.
-Reliability: Maintain stability and consistent functionality without frequent interruptions.
-Scalability: Design the bot to handle increasing server loads and accommodate future feature expansions.
-Security: Implement secure practices to prevent unauthorized access and data breaches.

Frameworks and Libraries:

-Go Lang: Primary programming language for bot development.
-discordgo: Library for interacting with Discord API.
-fmt, os, os/signal, strings, syscall, time: Go standard libraries used for various functionalities.
-github.com/bwmarrin/discordgo: Specifically for Discord integration.
-OS: Develop the bot to be compatible with Windows or Linux.

Programming Languages: Primarily using Go Lang.

Integrations:

-Discord API: Utilize Discord's official API for bot interactions.

Artifacts:

-Documentation: Comprehensive documentation detailing bot functionality, commands, and setup instructions.
-Test Cases: Develop test cases to ensure the bot operates as intended.
-User Manual: Create a user manual explaining how to use various bot features.
