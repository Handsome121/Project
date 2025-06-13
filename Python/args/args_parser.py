import argparse


def get_parser():
    parser = argparse.ArgumentParser(description='My Program')
    parser.add_argument('--input', type=str, help='input file path', required=True)
    parser.add_argument('--output', type=str, help='output file path', default='output.txt')
    parser.add_argument('--verbose', action='store_true', help='print verbose output')
    return parser


if __name__ == '__main__':
    args= get_parser().parse_known_args()
    print(args[0].__dict__)


