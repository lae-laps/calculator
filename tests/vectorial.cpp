#include <iostream>
#include <SFML/Graphics.hpp>

#define AVERAGE_BASE_OFFSET 10
#define AVERAGE_BASE_LENGTH 600

int main() {
    sf::ContextSettings settings;
    settings.antialiasingLevel = 16;
    sf::RenderWindow window(sf::VideoMode(800, 600), "main", sf::Style::Default, settings); 
    window.setVerticalSyncEnabled(true);
    
    // define the base for the vectors
    sf::VertexArray base(sf::LinesStrip, 3);

    base[0].position = sf::Vector2f(AVERAGE_BASE_OFFSET, AVERAGE_BASE_OFFSET);
    base[1].position = sf::Vector2f(AVERAGE_BASE_OFFSET, AVERAGE_BASE_LENGTH + AVERAGE_BASE_OFFSET);
    base[2].position = sf::Vector2f(AVERAGE_BASE_LENGTH + AVERAGE_BASE_OFFSET, AVERAGE_BASE_LENGTH + AVERAGE_BASE_OFFSET);

    // define the color of the base's points
    // base[0].color = sf::Color::Red;
    // base[1].color = sf::Color::Green;
    // base[2].color = sf::Color::Blue;

    sf::VertexArray vector1(sf::Lines, 2);

    vector1[0].position = sf::Vector2f(AVERAGE_BASE_OFFSET, AVERAGE_BASE_LENGTH + AVERAGE_BASE_OFFSET);
    vector1[1].position = sf::Vector2f(200, 200);

    vector1[0].color = sf::Color::Green;
    vector1[1].color = sf::Color::Green;

    while (window.isOpen()) {
        sf::Event event;
        while (window.pollEvent(event)) {
            if (event.type == sf::Event::Closed) {
                window.close();
            }
        }

        /*
        float x;
        std::cout << "x: " << std::endl;
        std::cin >> x;
        float y;
        std::cout << "y: " << std::endl;
        std::cin >> y;
        vector1[1].position = sf::Vector2f(x, y);
        */

        //vector1[1].position = sf::Vector2f(vector1[1].position.x - 1, vector1[1].position.y - 1);

        sf::Vector2i mouse = sf::Mouse::getPosition(window);
        vector1[1].position = sf::Vector2f(window.mapPixelToCoords(mouse));

        window.clear();
        window.draw(base);
        window.draw(vector1);
        window.display();
    }

    return 0;
}
