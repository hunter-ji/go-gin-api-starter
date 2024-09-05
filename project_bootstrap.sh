#!/bin/bash

OLD_PROJECT_NAME="Template"

BOLD='\033[1m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

checkSys=$(uname -s)

# welcome
echo -e "${BOLD}${BLUE}🚀 Welcome to the Go Project Initializer!${NC}"

# new project name
echo -e "\n${YELLOW}📝 Please enter the new project name:${NC}"
read -p "> " NEW_PROJECT_NAME

echo -e "\n${GREEN}Initializing your project...${NC}"

if [ "$checkSys" == "Darwin" ]; then
    # MacOS
    find . -type f \( -name "*.go" -o -name "go.mod" -o -name "*.md" -o -name "*.yaml" -o -name "*.yml" \) -exec sed -i "" "s/$OLD_PROJECT_NAME/$NEW_PROJECT_NAME/g" {} +
elif [ "$checkSys" == "Linux" ]; then
    # Linux
    find . -type f \( -name "*.go" -o -name "go.mod" -o -name "*.md" -o -name "*.yaml" -o -name "*.yml" \) -exec sed -i "s/$OLD_PROJECT_NAME/$NEW_PROJECT_NAME/g" {} +
else
    echo -e "${RED}Unsupported operating system.${NC}"
    exit 1
fi

echo -e "${GREEN}Project name updated successfully!${NC}"

# remove RabbitMQ folder if not needed
echo -e "\n${YELLOW}🐰 Do you need RabbitMQ in your project? [Y/n]:${NC}"
read -p "> " need_rabbitmq
need_rabbitmq=${need_rabbitmq:-Y}
if [[ $need_rabbitmq =~ ^[Nn]$ ]]; then
    if [ -d "pkg/util/rabbitMQ" ]; then
        rm -rf pkg/util/rabbitMQ
        echo -e "${GREEN}RabbitMQ folder removed.${NC}"
    else
        echo -e "${BLUE}RabbitMQ folder not found. No action taken.${NC}"
    fi
else
    echo -e "${BLUE}RabbitMQ folder kept intact.${NC}"
fi

# remove .git directory if needed
echo -e "\n${YELLOW}🗑️ Do you want to remove the existing .git directory? [Y/n]:${NC}"
read -p "> " remove_git
remove_git=${remove_git:-Y}
if [[ $remove_git =~ ^[Yy]$ ]]; then
    rm -rf .git
    echo -e "${GREEN}Existing .git directory removed.${NC}"

    # new Git repository
    echo -e "\n${YELLOW}🔧 Do you want to initialize a new Git repository? [Y/n]:${NC}"
    read -p "> " init_new_git
    init_new_git=${init_new_git:-Y}
    if [[ $init_new_git =~ ^[Yy]$ ]]; then
        git init
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}New Git repository initialized successfully.${NC}"
        else
            echo -e "${RED}Failed to initialize new Git repository.${NC}"
        fi
    else
        echo -e "${BLUE}Skipped initializing new Git repository.${NC}"
    fi
else
    echo -e "${BLUE}Existing .git directory kept intact.${NC}"
fi

echo -e "\n${BOLD}${GREEN}🎉 Project initialization complete!${NC}"
echo -e "${BLUE}Summary:${NC}"
echo -e "   ${YELLOW}Old project name:${NC} $OLD_PROJECT_NAME"
echo -e "   ${YELLOW}New project name:${NC} $NEW_PROJECT_NAME"
echo -e "   ${YELLOW}RabbitMQ:${NC} $([[ $need_rabbitmq =~ ^[Yy]$ ]] && echo "Included" || echo "Removed")"
echo -e "   ${YELLOW}Git:${NC} $([[ $remove_git =~ ^[Yy]$ ]] && ([[ $init_new_git =~ ^[Yy]$ ]] && echo "Reinitialized" || echo "Removed") || echo "Kept existing")"

# go mod tidy
echo -e "\n${YELLOW}🧹 Do you want to run 'go mod tidy' to clean up dependencies? [Y/n]:${NC}"
read -p "> " run_tidy
run_tidy=${run_tidy:-Y}
if [[ $run_tidy =~ ^[Yy]$ ]]; then
    echo -e "${BLUE}Running 'go mod tidy'...${NC}"
    go mod tidy
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Dependencies cleaned up successfully.${NC}"
    else
        echo -e "${RED}Error occurred while cleaning up dependencies.${NC}"
    fi
else
    echo -e "${BLUE}Skipped 'go mod tidy'. Remember to run it later if needed.${NC}"
fi

# new steps
echo -e "\n${BOLD}${CYAN}🚀 How to start your project:${NC}"
echo -e "${CYAN}1. Modify .env.development in the project root to set Redis and PostgreSQL configurations.${NC}"
echo -e "${CYAN}2. Run the following command to start your project:${NC}"
echo -e "   ${YELLOW}make run${NC}"

echo -e "\n${BOLD}${GREEN}🎈 Your project is ready! Happy coding!${NC}"
