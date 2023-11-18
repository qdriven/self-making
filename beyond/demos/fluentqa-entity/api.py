import inspect
import steamship

from steamship import check_environment, RuntimeEnvironments, Steamship
from steamship.invocable import post, PackageService

class EntityGeneratePackage(PackageService):
    PROMPT = "根据以下内容:{references},生成{language}基于{lib_name}实体类"

    @post("generate")
    def generate(self, language: str, lib_name: str, references: str):
        gpt4 = self.client.use_plugin("gpt-4")
        task = gpt4.generate(text=self.PROMPT.format(language=language, 
                    lib_name=lib_name, references=references))
        task.wait()
        return task.output.blocks[0].text


handler = steamship.invocable.create_handler(EntityGeneratePackage)

if __name__ == '__main__':
    from termcolor import colored
    print(colored("Generate Response By GPT-4\n", attrs=['bold']))
    # This helper provides runtime API key prompting, etc.
    check_environment(RuntimeEnvironments.REPLIT)
    with Steamship.temporary_workspace() as client:
        prompt = EntityGeneratePackage(client=client)
        print(colored("Now, try with your inputs...", 'green'))
        try_again = True
        while try_again:
            kwargs = {}
            for parameter in inspect.signature(prompt.generate).parameters:
                kwargs[parameter] = input(
                    colored(f'{parameter.capitalize()}: ', 'grey'))
            print(colored("Generating...", 'grey'))
            print(colored("Compliment:", 'grey'), f'{prompt.generate(**kwargs)}\n')

            try_again = input(colored("Generate another (y/n)? ",
                                      'green')).lower().strip() == 'y'
            print()
            print("Ready to share with your friends (and the world)?")
            print("Run ", colored("$ ship deploy ", color='green',
                                  on_color='on_black'),
                  "to get a production-ready API endpoint and web-based demo app.")