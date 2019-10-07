import color.Color;
import color.Red;
import shape.Circle;
import shape.Shape;

/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 13:43
 **/
public class RedCircleFactory extends AbstractFactory {
    @Override
    public Shape getShape() {
        return new Circle();
    }

    @Override
    public Color getColor() {
        return new Red();
    }
}
